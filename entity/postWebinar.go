package entity

import (
	"time"

	"github.com/google/uuid"
)

type PostWebinar struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	UserId      uuid.UUID `json:"userId" gorm:"foreignkey:ID;references:users"`
	WebinarName string    `json:"name" gorm:"not null; type:varchar(255)"`
	Description string    `json:"description" gorm:"type:text"`
	Price       string    `json:"price" gorm:""`
	EventDate   time.Time `json:"eventDate"`
	IsBought    bool      `json:"isBought"`
}
