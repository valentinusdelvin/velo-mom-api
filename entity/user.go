package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
	Email       string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password    string    `json:"password" gorm:"type:varchar(255);not null"`
	DisplayName string    `json:"displayName" gorm:"type:varchar(255);not null"`
	PhoneNumber string    `json:"phoneNumber"`
	Bio         string    `json:"bio"`
	IsAdmin     bool      `json:"isAdmin"`
}
