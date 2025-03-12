package entity

import (
	"time"

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

type Article struct {
	ID            uuid.UUID `json:"id" gorm:"primaryKey;unique;"`
	Title         string    `json:"title" gorm:"not null"`
	Content       string    `json:"content" gorm:"not null"`
	Summary       string    `json:"summary" gorm:"not null"`
	Author        string    `json:"author" gorm:"not null"`
	CreatedAt     string    `json:"createdAt"`
	Def_CreatedAt time.Time `gorm:"autoCreateTime"`
	ImageURL      string    `json:"imageURL"`
	Filter        Filter    `gorm:"filter"`
}
