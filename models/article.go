package models

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type Filter int

const (
	BabyBlues         Filter = 1
	SelfCare          Filter = 2
	KesehatanMental   Filter = 3
	Journaling        Filter = 4
	DukunganEmosional Filter = 5
	Keseharian        Filter = 6
	TipsRelaksasi     Filter = 7
	Parenting         Filter = 8
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
	Filter    Filter                `form:"filter"`
}

type GetArticles struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	CreatedAt string    `json:"createdAt"`
	ImageURL  string    `json:"imageURL"`
	Filter    Filter    `json:"filter"`
}
