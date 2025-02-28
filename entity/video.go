package entity

import "time"

type Video struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Tittle      string    `json:"tittle" gorm:"not null"`
	Link        string    `json:"link" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
