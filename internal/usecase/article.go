package usecase

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"github.com/valentinusdelvin/velo-mom-api/pkg/supabase"
	addition "github.com/valentinusdelvin/velo-mom-api/pkg/timeconvert"
)

type InterArticleUsecase interface {
	CreateArticle(param models.CreateArticle) error
	GetArticles(page, size int) ([]models.GetArticles, error)
	GetArticleByID(id string) (entity.Article, error)
	GetArticlesBySearch(param models.GetArticles, page, size int) ([]models.GetArticles, error)
	GetArticleByFilter(param models.GetArticles, page, size int) ([]models.GetArticles, error)
}

type ArticleUsecase struct {
	arsc repository.InterArticleRepository
	sb   supabase.InterSupabase
}

func NewArticleUsecase(articleRepo repository.InterArticleRepository, supabase supabase.InterSupabase) InterArticleUsecase {
	return &ArticleUsecase{
		arsc: articleRepo,
		sb:   supabase,
	}
}

func (a *ArticleUsecase) CreateArticle(param models.CreateArticle) error {
	param.ID = uuid.New()
	ext := filepath.Ext(param.PhotoIMG.Filename)
	if ext == "" {
		return errors.New("invalid file extension: no file extension found")
	}
	param.PhotoIMG.Filename = fmt.Sprintf("%s-%v%s", param.ID, time.Now().Unix(), ext)

	newPhotoLink, err := a.sb.Upload(param.PhotoIMG)
	if err != nil {
		return err
	}

	articlePost := entity.Article{
		ID:            param.ID,
		Title:         param.Title,
		Content:       param.Content,
		Summary:       param.Summary,
		Author:        param.Author,
		ImageURL:      newPhotoLink,
		Def_CreatedAt: time.Now(),
		CreatedAt:     addition.TimeConvert(time.Now()),
	}

	_, err = a.arsc.CreateArticle(articlePost)
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleUsecase) GetArticles(page, size int) ([]models.GetArticles, error) {
	articles, err := a.arsc.GetArticles(page, size)
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

func (a *ArticleUsecase) GetArticlesBySearch(param models.GetArticles, page, size int) ([]models.GetArticles, error) {
	articles, err := a.arsc.GetArticlesBySearch(param, page, size)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleUsecase) GetArticleByFilter(param models.GetArticles, page, size int) ([]models.GetArticles, error) {
	articles, err := a.arsc.GetArticleByFilter(param, page, size)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
