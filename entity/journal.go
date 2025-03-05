package entity

import "github.com/google/uuid"

type Emoji int

const (
	EmojiTired     Emoji = 1
	EmojiSad       Emoji = 2
	EmojiNeutral   Emoji = 3
	EmojiHappy     Emoji = 4
	EmojiVeryHappy Emoji = 5
)

type Journal struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"foreignkey:ID;references:users"`
	CreatedAt string    `json:"created_at"`
	Title     string    `gorm:"not null"`
	Story     string
	Feels     string
	Emoji     Emoji `gorm:"emoji"`
}
