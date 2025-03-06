package repository

import (
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"gorm.io/gorm"
)

type InterUserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(param models.UserParam) (entity.User, error)
	UpdateUser(param models.UserUpdate, user entity.User) error
	UpdateProfilePhoto(param models.UpdateProfilePhoto, user entity.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) InterUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) GetUser(param models.UserParam) (entity.User, error) {
	user := entity.User{}
	err := ur.db.Where(&param).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateUser(param models.UserUpdate, user entity.User) error {
	err := ur.db.Model(&user).Updates(param).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UpdateProfilePhoto(param models.UpdateProfilePhoto, user entity.User) error {
	err := ur.db.Model(&user).Where("id = ?", param.ID).Update("photo_link", param.PhotoLink).Error
	if err != nil {
		return err
	}

	return nil
}
