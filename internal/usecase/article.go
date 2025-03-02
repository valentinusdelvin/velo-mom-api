package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/models"
	addition "github.com/valentinusdelvin/velo-mom-api/utils/timeconvert"
)

type InterArticleUsecase interface {
	CreateArticle(param models.CreateArticle) error
	GetArticles() ([]models.GetArticles, error)
	GetArticleByID(id string) (entity.Article, error)
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
	articlePost := entity.Article{
		ID:        uuid.New(),
		Title:     param.Title,
		Content:   param.Content,
		Summary:   param.Summary,
		Author:    param.Author,
		ImageURL:  param.ImageURL,
		CreatedAt: addition.TimeConvert(time.Now()),
	}

	_, err := a.arsc.CreateArticle(articlePost)
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleUsecase) GetArticles() ([]models.GetArticles, error) {
	articles, err := a.arsc.GetArticles()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleUsecase) GetArticleByID(id string) (entity.Article, error) {
	article, err := a.arsc.GetArticleByID(id)
	if err != nil {
		return entity.Article{}, err
	}
	return article, nil
}
