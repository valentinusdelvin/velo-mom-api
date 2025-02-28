package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateArticle struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title" binding:"required,min=3"`
	Content   string    `json:"content" binding:"required,min=10"`
	Author    string    `json:"author" binding:"required,min=1"`
	ImageURL  string    `json:"imageURL"`
	CreatedAt time.Time `json:"createdAt"`
}
