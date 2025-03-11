package models

import "github.com/google/uuid"

type Payment struct {
	OrderID   uuid.UUID
	UserID    uuid.UUID
	ProductID uuid.UUID
}
