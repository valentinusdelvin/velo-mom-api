package rest

import (
	"net/http"
	"strconv"

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
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "9"))

	videos, err := r.usecase.VideoUsecase.GetVideos(page, size)
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

func (r *Rest) GetVideosBySearch(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "9"))

	keyword := ctx.Query("keyword")
	if keyword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad keyword"})
		return
	}

	param := models.CreateVideo{
		Title: keyword,
	}
	videos, err := r.usecase.VideoUsecase.GetVideosBySearch(param, page, size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get videos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"videos": videos})
}

func ConvertInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func (r *Rest) GetVideoByFilter(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "9"))

	category := ctx.Query("category")
	if category == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad category"})
		return
	}

	int_category, err := ConvertInt(category)
	if err != nil {
		return
	}

	param := models.CreateVideo{
		Filter: models.Filter(int_category),
	}
	videos, err := r.usecase.VideoUsecase.GetVideoByFilter(param, page, size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get videos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"videos": videos})
}
