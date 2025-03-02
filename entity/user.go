package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	Email       string    `gorm:"type:varchar(255);unique;not null"`
	Password    string    `gorm:"type:varchar(255);not null"`
	DisplayName string    `gorm:"type:varchar(255);not null"`
	PhoneNumber string
	Bio         string
	IsAdmin     bool
}
