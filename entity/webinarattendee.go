package entity

import "github.com/google/uuid"

type WebinarAttendee struct {
	WebinarID uuid.UUID
	UserID    uuid.UUID

	Webinar Webinar `gorm:"foreignKey:WebinarID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User    User    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
