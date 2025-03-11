package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
)

func (r *Rest) Purchase(ctx *gin.Context) {
	user, ok := ctx.Get("userID")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "failed to get user",
		})
		return
	}

	parsedID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to parse product ID",
		})
		return
	}

	payment := entity.Payment{
		OrderID:   uuid.New(),
		UserID:    user.(uuid.UUID),
		ProductID: parsedID,
	}

	paymentLink, err := r.usecase.PaymentUsecase.Purchase(payment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to Purchase",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"payment_link": paymentLink})
}

func (r *Rest) Validate(ctx *gin.Context) {
	var MidtransNotifications map[string]interface{}

	err := ctx.ShouldBind(&MidtransNotifications)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Notification Payload"})
	}

	_, exists := MidtransNotifications["order_id"].(string)
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Missing order_id in notification"})
	}

	err = r.usecase.PaymentUsecase.Validate(MidtransNotifications)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to validate payment",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Payment validated successfully",
	})
}
