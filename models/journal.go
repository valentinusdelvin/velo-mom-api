package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateJournal struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	Title         string `json:"title" binding:"required"`
	Story         string `json:"story"`
	Feels         string `json:"feels"`
	Emoji         int    `json:"emoji"`
	Def_CreatedAt time.Time
	CreatedAt     string `json:"created_at"`
}
