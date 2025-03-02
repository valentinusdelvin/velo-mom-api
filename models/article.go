package models

import (
	"github.com/google/uuid"
)

type CreateArticle struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title" binding:"required,min=3"`
	Content   string    `json:"content" binding:"required,min=10"`
	Summary   string    `json:"summary" binding:"required,min=5"`
	Author    string    `json:"author" binding:"required,min=1"`
	ImageURL  string    `json:"imageURL"`
	CreatedAt string    `json:"createdAt"`
}

type GetArticles struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title" binding:"required,min=3"`
	Summary   string    `json:"summary" binding:"required,min=5"`
	CreatedAt string    `json:"createdAt"`
	ImageURL  string    `json:"imageURL"`
}
