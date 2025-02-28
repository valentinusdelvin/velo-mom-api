package repository

import (
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"gorm.io/gorm"
)

type InterArticleRepository interface {
	CreateArticle(article entity.Article) (entity.Article, error)
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
