package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
	Username    string    `json:"username" gorm:"type:varchar(255);unique;not null"`
	Email       string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password    string    `json:"password" gorm:"type:varchar(255);not null"`
	DisplayName string    `json:"displayName" gorm:"type:varchar(255);not null"`
	Age         uint      `json:"age"`
	PhoneNumber string    `json:"phoneNumber"`
	IsAdmin     bool      `json:"isAdmin"`
}
