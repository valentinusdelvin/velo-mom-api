package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/models"
)

func (r *Rest) CreateJournal(ctx *gin.Context) {
	param := models.CreateJournal{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind with JSON",
			"error":   err,
		})
		return
	}

	user, ok := ctx.Get("userID")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to get login user",
		})
		return
	}

	param.UserID = user.(uuid.UUID)

	err = r.usecase.JournalUsecase.CreateJournal(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create journal",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success to create journal"})
}

func (r *Rest) GetUserJournals(ctx *gin.Context) {
	user, authorized := ctx.Get("userID")
	if !authorized {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userID, ok := user.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id"})
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "4"))
	journals, err := r.usecase.JournalUsecase.GetUserJournals(userID, page, size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve journals"})
		return
	}
	ctx.JSON(http.StatusOK, journals)
}

func (r *Rest) GetUserJournalByID(ctx *gin.Context) {
	user, authorized := ctx.Get("userID")
	if !authorized {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userID, ok := user.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id"})
		return
	}

	id := ctx.Param("id")

	journal, err := r.usecase.JournalUsecase.GetUserJournalByID(userID, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get journal",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, journal)
}
