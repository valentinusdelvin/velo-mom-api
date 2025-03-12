package repository

import (
	"fmt"

	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"gorm.io/gorm"
)

type InterArticleRepository interface {
	CreateArticle(article entity.Article) (entity.Article, error)
	GetArticles(page, size int) ([]models.GetArticles, error)
	GetArticleByID(id string) (entity.Article, error)
	GetArticlesBySearch(param models.GetArticles, page, size int) ([]models.GetArticles, error)
	GetArticleByFilter(param models.GetArticles, page, size int) ([]models.GetArticles, error)
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

func (ar *ArticleRepository) GetArticles(page, size int) ([]models.GetArticles, error) {
	var articles []models.GetArticles

	offset := (page - 1) * size

	err := ar.db.Table("articles").Order("Def_CreatedAt DESC").Limit(size).Offset(offset).Find(&articles).Error
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

func (ar *ArticleRepository) GetArticlesBySearch(param models.GetArticles, page, size int) ([]models.GetArticles, error) {
	var articles []models.GetArticles

	offset := (page - 1) * size

	err := ar.db.Table("articles").Order("Def_CreatedAt DESC").Limit(size).Offset(offset).Where("title ILIKE ?", fmt.Sprintf("%%%s%%", param.Title)).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (ar *ArticleRepository) GetArticleByFilter(param models.GetArticles, page, size int) ([]models.GetArticles, error) {
	var articles []models.GetArticles

	offset := (page - 1) * size

	err := ar.db.Model(entity.Article{}).Order("Def_CreatedAt DESC").Limit(size).Offset(offset).Where("filter = ?", param.Filter).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}
