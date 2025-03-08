package repository

import (
	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"gorm.io/gorm"
)

type InterUserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(param models.UserParam) (entity.User, error)
	GetUserInfo(id uuid.UUID) (entity.User, error)
	UpdateUser(param models.UserUpdate, id uuid.UUID) error
	UpdateProfilePhoto(param models.UpdateProfilePhoto, id uuid.UUID) error
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

func (ur *UserRepository) GetUserInfo(id uuid.UUID) (entity.User, error) {
	user := entity.User{}
	err := ur.db.Select("id, display_name, email, phone_number, bio, photo_link").Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateUser(param models.UserUpdate, id uuid.UUID) error {
	err := ur.db.Model(&entity.User{}).Where("id = ?", id).Updates(param).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UpdateProfilePhoto(param models.UpdateProfilePhoto, id uuid.UUID) error {
	err := ur.db.Model(&entity.User{}).Where("id = ?", id).Update("photo_link", param.PhotoLink).Error
	if err != nil {
		return err
	}

	return nil
}
