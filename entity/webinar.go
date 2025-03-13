package entity

import (
	"time"

	"github.com/google/uuid"
)

type Webinar struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	WebinarName string    `gorm:"not null; type:varchar(255)"`
	Subheader   string    `gorm:"not null; type:text"`
	Description string    `gorm:"type:text"`
	Price       uint64    `gorm:"not null"`
	Photolink   string
	Quota       uint `gorm:"not null"`
	EventDate   time.Time
	StrDate     string
	EventTime   string `gorm:"not null"`
	Location    string `gorm:"not null"`
	CreatedAt   time.Time

	WebinarAttendees []WebinarAttendee `gorm:"foreignKey:WebinarID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
