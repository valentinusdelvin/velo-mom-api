package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/models"
)

type InterArticleUsecase interface {
	CreateArticle(param models.CreateArticle) error
}

type ArticleUsecase struct {
	arsc repository.InterArticleRepository
}

func NewArticleUsecase(articleRepo repository.InterArticleRepository) InterArticleUsecase {
	return &ArticleUsecase{
		arsc: articleRepo,
	}
}

func (a *ArticleUsecase) CreateArticle(param models.CreateArticle) error {
	if len(param.Title) < 3 {
		return errors.New("title must be contain 3 or more characters")
	}
	if len(param.Content) < 10 {
		return errors.New("content must be contain 10 or more characters")
	}
	if len(param.Author) < 1 {
		return errors.New("author must be contain 1 or more characters")
	}

	articlePost := entity.Article{
		ID:        uuid.New(),
		Title:     param.Title,
		Content:   param.Content,
		Author:    param.Author,
		ImageURL:  param.ImageURL,
		CreatedAt: time.Now(),
	}

	_, err := a.arsc.CreateArticle(articlePost)
	if err != nil {
		return err
	}
	return nil
}
