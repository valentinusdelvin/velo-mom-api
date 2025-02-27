package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository InterUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	UserRepository := NewUserRepository(db)

	return &Repository{
		UserRepository: UserRepository,
	}
}
