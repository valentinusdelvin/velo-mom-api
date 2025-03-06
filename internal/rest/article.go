package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/valentinusdelvin/velo-mom-api/models"
)

func (r *Rest) CreateArticle(ctx *gin.Context) {
	param := models.CreateArticle{}

	err := ctx.ShouldBindWith(&param, binding.FormMultipart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request body",
			"error":   err,
		})
		return
	}

	err = r.usecase.ArticleUsecase.CreateArticle(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create article",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success to create article"})
}

func (r *Rest) GetArticles(ctx *gin.Context) {
	articles, err := r.usecase.ArticleUsecase.GetArticles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get articles",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, articles)
}

func (r *Rest) GetArticleByID(ctx *gin.Context) {
	id := ctx.Param("id")

	article, err := r.usecase.ArticleUsecase.GetArticleByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get article",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, article)
}
