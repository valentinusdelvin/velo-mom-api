package models

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type CreateArticle struct {
	ID        uuid.UUID             `json:"id"`
	Title     string                `form:"title" binding:"required,min=3"`
	Content   string                `form:"content" binding:"required,min=10"`
	Summary   string                `form:"summary" binding:"required,min=5"`
	Author    string                `form:"author" binding:"required,min=1"`
	ImageURL  string                `form:"imageURL"`
	PhotoIMG  *multipart.FileHeader `form:"photo" binding:"required"`
	CreatedAt string                `json:"createdAt"`
}

type GetArticles struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	CreatedAt string    `json:"createdAt"`
	ImageURL  string    `json:"imageURL"`
}
