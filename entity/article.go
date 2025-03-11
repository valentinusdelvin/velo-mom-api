package entity

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID            uuid.UUID `json:"id" gorm:"primaryKey;unique;"`
	Title         string    `json:"title" gorm:"not null"`
	Content       string    `json:"content" gorm:"not null"`
	Summary       string    `json:"summary" gorm:"not null"`
	Author        string    `json:"author" gorm:"not null"`
	CreatedAt     string    `json:"createdAt"`
	Def_CreatedAt time.Time `gorm:"autoCreateTime"`
	ImageURL      string    `json:"imageURL"`
}
