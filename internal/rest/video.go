package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/velo-mom-api/models"
)

func (r *Rest) CreateVideo(ctx *gin.Context) {
	param := models.CreateVideo{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind with JSON",
			"error":   err,
		})
		return
	}

	err = r.usecase.VideoUsecase.CreateVideo(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create videos",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success to create video"})
}

func (r *Rest) GetVideos(ctx *gin.Context) {
	videos, err := r.usecase.VideoUsecase.GetVideos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get articles",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, videos)
}

func (r *Rest) GetVideoByID(ctx *gin.Context) {
	id := ctx.Param("id")

	video, err := r.usecase.VideoUsecase.GetVideoByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get webinar",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, video)
}
