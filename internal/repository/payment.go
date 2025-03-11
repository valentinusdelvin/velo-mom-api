package repository

import (
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"gorm.io/gorm"
)

type InterPaymentRepository interface {
	CreatePayment(payment entity.Payment) error
	UpdatePaymentStatus(tx *gorm.DB, status string, orderID string) error
	GetInvoice(orderID string) (models.Payment, error)
}

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) InterPaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (pr *PaymentRepository) CreatePayment(payment entity.Payment) error {
	err := pr.db.Create(&payment).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *PaymentRepository) UpdatePaymentStatus(tx *gorm.DB, status string, orderID string) error {
	return tx.Model(&entity.Payment{}).
		Where("order_id = ?", orderID).
		Update("status", status).
		Error
}

func (pr *PaymentRepository) GetInvoice(orderID string) (models.Payment, error) {
	var invoice models.Payment
	err := pr.db.Table("payments").Where("order_id = ?", orderID).First(&invoice).Error
	return invoice, err
}
