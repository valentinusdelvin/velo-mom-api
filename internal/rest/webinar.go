package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/valentinusdelvin/velo-mom-api/models"
)

func (r *Rest) CreateWebinar(ctx *gin.Context) {
	param := models.CreateWebinar{}

	err := ctx.ShouldBindWith(&param, binding.FormMultipart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request body",
			"error":   err.Error(),
		})
		return
	}

	err = r.usecase.WebinarUsecase.CreateWebinar(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create webinar",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success to create webinar"})
}

func (r *Rest) GetWebinars(ctx *gin.Context) {
	webinars, err := r.usecase.WebinarUsecase.GetWebinars()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get webinars",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, webinars)
}

func (r *Rest) GetWebinarByID(ctx *gin.Context) {
	id := ctx.Param("id")

	webinar, err := r.usecase.WebinarUsecase.GetWebinarByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get webinar",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, webinar)
}
