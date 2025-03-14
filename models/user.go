package models

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UserRegister struct {
	ID          uuid.UUID `json:"-"`
	DisplayName string    `json:"name" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Password    string    `json:"password" binding:"required,min=6"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserInfo struct {
	PhotoLink   string `json:""`
	DisplayName string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Bio         string `json:"bio"`
}

type UserParam struct {
	ID          uuid.UUID `json:"-"`
	DisplayName string    `json:"-"`
	Email       string    `json:"-"`
}

type EmailAuthenticator struct {
	Email string
}

type UserUpdate struct {
	DisplayName string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Bio         string `json:"bio"`
}

type UpdateProfilePhoto struct {
	ID        uuid.UUID             `json:"-"`
	PhotoLink string                `json:"-"`
	PhotoIMG  *multipart.FileHeader `form:"photo" binding:"required"`
}
