package models

import (
	"github.com/google/uuid"
)

type UserRegister struct {
	ID          uuid.UUID `json:"-"`
	DisplayName string    `json:"displayName" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Password    string    `json:"password" binding:"required,min=6"`
}

type UserLogin struct {
	DisplayName string `json:"displayName" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
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
	DisplayName string `json:"displayName"`
	PhoneNumber string `json:"phoneNumber"`
	Bio         string `json:"bio"`
}
