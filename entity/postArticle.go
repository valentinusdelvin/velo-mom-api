package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type PostArticle struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;unique;"`
	Title     string         `json:"title" gorm:"not null"`
	Content   string         `json:"content" gorm:"not null"`
	Author    string         `json:"author"`
	Category  pq.StringArray `json:"category"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	ImageURL  string         `json:"imageURL"`
}
