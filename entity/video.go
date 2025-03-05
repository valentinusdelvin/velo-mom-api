package entity

import "time"

type Video struct {
	YoutubeID   string `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	YoutubeURL  string `gorm:"not null"`
	Description string
	Thumbnail   string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
