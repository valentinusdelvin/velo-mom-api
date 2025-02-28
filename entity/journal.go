package entity

import "github.com/google/uuid"

type Journal struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
}
