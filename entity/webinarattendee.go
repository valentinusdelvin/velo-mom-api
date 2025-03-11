package entity

import "github.com/google/uuid"

type WebinarAttendee struct {
	WebinarID uuid.UUID
	UserID    uuid.UUID
}
