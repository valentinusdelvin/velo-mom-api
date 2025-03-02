package entity

import "time"

type Video struct {
	ID          string    `gorm:"primaryKey"`
	Tittle      string    `gorm:"not null"`
	Link        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
