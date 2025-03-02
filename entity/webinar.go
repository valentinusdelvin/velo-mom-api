package entity

import (
	"time"

	"github.com/google/uuid"
)

type PostWebinar struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	WebinarName string    `gorm:"not null; type:varchar(255)"`
	Description string    `gorm:"type:text"`
	Price       string    `gorm:""`
	EventDate   time.Time `gorm:"not null"`
	IsBought    bool      `gorm:"default:false"`
}
