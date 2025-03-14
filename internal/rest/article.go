package rest

import (
	"net/http"
	"strconv"

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

func (r *Rest) DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	err := r.usecase.ArticleUsecase.DeleteArticle(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "article deleted successfully"})
}

func (r *Rest) GetArticles(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "9"))
	articles, err := r.usecase.ArticleUsecase.GetArticles(page, size)
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

func (r *Rest) GetArticlesBySearch(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "9"))

	keyword := ctx.Query("keyword")
	if keyword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad keyword"})
		return
	}

	param := models.GetArticles{
		Title: keyword,
	}
	articles, err := r.usecase.ArticleUsecase.GetArticlesBySearch(param, page, size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get articles"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"articles": articles})

}

func (r *Rest) GetArticleByFilter(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "9"))

	filter := ctx.Query("filter")
	if filter == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad category"})
		return
	}

	int_category, err := ConvertInt(filter)
	if err != nil {
		return
	}

	param := models.GetArticles{
		Filter: models.Filter(int_category),
	}
	articles, err := r.usecase.ArticleUsecase.GetArticleByFilter(param, page, size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get articles"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"articles": articles})
}
