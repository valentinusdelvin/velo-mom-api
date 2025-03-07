package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository    InterUserRepository
	ArticleRepository InterArticleRepository
	VideoRepository   InterVideoRepository
	JournalRepository InterJournalRepository
	WebinarRepository InterWebinarRepository
}

func NewRepository(db *gorm.DB) *Repository {
	UserRepository := NewUserRepository(db)
	ArticleRepository := NewArticleRepository(db)
	VideoRepository := NewVideoRepository(db)
	JournalRepository := NewJournalRepository(db)
	WebinarRepository := NewWebinarRepository(db)

	return &Repository{
		UserRepository:    UserRepository,
		ArticleRepository: ArticleRepository,
		VideoRepository:   VideoRepository,
		JournalRepository: JournalRepository,
		WebinarRepository: WebinarRepository,
	}
}
