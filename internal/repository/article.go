package repository

import (
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"gorm.io/gorm"
)

type InterArticleRepository interface {
	CreateArticle(article entity.Article) (entity.Article, error)
	GetArticles() ([]models.GetArticles, error)
	GetArticleByID(id string) (entity.Article, error)
}

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) InterArticleRepository {
	return &ArticleRepository{
		db: db,
	}
}

func (ar *ArticleRepository) CreateArticle(article entity.Article) (entity.Article, error) {
	err := ar.db.Create(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

func (ar *ArticleRepository) GetArticles() ([]models.GetArticles, error) {
	var articles []models.GetArticles

	err := ar.db.Table("articles").Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (ar *ArticleRepository) GetArticleByID(id string) (entity.Article, error) {
	var article entity.Article

	err := ar.db.Table("articles").Where("id = ?", id).Find(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}
