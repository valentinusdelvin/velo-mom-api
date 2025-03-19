package usecase

import (
	"errors"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"gorm.io/gorm"
)

type InterPaymentUsecase interface {
	Purchase(payment entity.Payment) (string, error)
	Validate(MidtransNotifications map[string]interface{}) error
}

type PaymentUsecase struct {
	prsc repository.InterPaymentRepository
	wrsc repository.InterWebinarRepository
	db   *gorm.DB
}

func NewPaymentUsecase(paymentRepo repository.InterPaymentRepository, webinarRepo repository.InterWebinarRepository, db *gorm.DB) InterPaymentUsecase {
	return &PaymentUsecase{
		prsc: paymentRepo,
		wrsc: webinarRepo,
		db:   db,
	}
}

func (p *PaymentUsecase) Purchase(payment entity.Payment) (string, error) {
	webinar, err := p.wrsc.GetWebinarByID(payment.ProductID.String())
	if err != nil {
		return "", err
	}
	if webinar.Quota <= 0 {
		return "", errors.New("webinar quota full")
	}

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payment.OrderID.String(),
			GrossAmt: int64(webinar.Price),
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:           webinar.ID.String(),
				Name:         webinar.WebinarName,
				Price:        int64(webinar.Price),
				Qty:          1,
				Category:     "Webinar VeloVent",
				MerchantName: "Velo Mom",
			},
		},
	}

	ServiceFee := midtrans.ItemDetails{
		ID:    "biaya_layanan",
		Name:  "Biaya Layanan",
		Price: 2000,
		Qty:   1,
	}

	*snapReq.Items = append(*snapReq.Items, ServiceFee)
	payment.Price = uint64(snapReq.TransactionDetails.GrossAmt)

	snapReq.TransactionDetails.GrossAmt += ServiceFee.Price
	payment.FinalPrice = uint64(snapReq.TransactionDetails.GrossAmt)

	paymentLink, paymentErr := snap.CreateTransactionUrl(snapReq)
	if paymentErr != nil {
		return "", paymentErr
	}
	payment.PaymentLink = paymentLink
	payment.ProductName = (*snapReq.Items)[0].Name

	err = p.prsc.CreatePayment(payment)
	if err != nil {
		return "", err
	}
	return paymentLink, nil
}

func (p *PaymentUsecase) Validate(MidtransNotifications map[string]interface{}) error {
	transactionStatus := MidtransNotifications["transaction_status"]
	orderID := MidtransNotifications["order_id"].(string)
	fraudStatus := MidtransNotifications["fraud_status"]

	err := p.db.Transaction(func(tx *gorm.DB) error {
		switch transactionStatus {
		case "capture":
			switch fraudStatus {
			case "challenge":
				return p.prsc.UpdatePaymentStatus(tx, "challenge", orderID)
			case "accept":
				if err := p.prsc.UpdatePaymentStatus(tx, "success", orderID); err != nil {
					return err
				}
				invoice, err := p.prsc.GetInvoice(orderID)
				if err != nil {
					return err
				}
				attendee := entity.WebinarAttendee{
					UserID:    invoice.UserID,
					WebinarID: invoice.ProductID,
				}

				if err := p.wrsc.CreateWebinarAttendee(tx, attendee); err != nil {
					return err
				}

				if err := p.wrsc.UpdateWebinarInfo(tx, invoice.ProductID); err != nil {
					return err
				}
				return nil
			}

		case "settlement":
			if err := p.prsc.UpdatePaymentStatus(tx, "success", orderID); err != nil {
				return err
			}
			invoice, err := p.prsc.GetInvoice(orderID)
			if err != nil {
				return err
			}
			attendee := entity.WebinarAttendee{
				UserID:    invoice.UserID,
				WebinarID: invoice.ProductID,
			}

			if err := p.wrsc.CreateWebinarAttendee(tx, attendee); err != nil {
				return err
			}

			if err := p.wrsc.UpdateWebinarInfo(tx, invoice.ProductID); err != nil {
				return err
			}
			return nil

		case "cancel", "expire":
			return p.prsc.UpdatePaymentStatus(tx, "failure", orderID)

		case "pending":
			return p.prsc.UpdatePaymentStatus(tx, "pending", orderID)

		case "deny":
			return p.prsc.UpdatePaymentStatus(tx, "denied", orderID)
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
