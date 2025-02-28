package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository    InterUserRepository
	ArticleRepository InterArticleRepository
}

func NewRepository(db *gorm.DB) *Repository {
	UserRepository := NewUserRepository(db)
	ArticleRepository := NewArticleRepository(db)

	return &Repository{
		UserRepository:    UserRepository,
		ArticleRepository: ArticleRepository,
	}
}
