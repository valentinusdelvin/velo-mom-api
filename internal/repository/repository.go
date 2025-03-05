package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository    InterUserRepository
	ArticleRepository InterArticleRepository
	VideoRepository   InterVideoRepository
	JournalRepository InterJournalRepository
}

func NewRepository(db *gorm.DB) *Repository {
	UserRepository := NewUserRepository(db)
	ArticleRepository := NewArticleRepository(db)
	VideoRepository := NewVideoRepository(db)
	JournalRepository := NewJournalRepository(db)

	return &Repository{
		UserRepository:    UserRepository,
		ArticleRepository: ArticleRepository,
		VideoRepository:   VideoRepository,
		JournalRepository: JournalRepository,
	}
}
