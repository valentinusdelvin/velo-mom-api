package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/velo-mom-api/models"
)

func (r *Rest) CreateArticle(ctx *gin.Context) {
	param := models.CreateArticle{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind with JSON",
			"error":   err,
		})
		return
	}

	err = r.usecase.ArticleUsecase.CreateArticle(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create article",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success to create article"})
}
