package models

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type CreateWebinar struct {
	ID          uuid.UUID             `form:"id"`
	WebinarName string                `form:"name" binding:"required"`
	Subheader   string                `form:"subheader" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Price       uint64                `form:"price" binding:"required"`
	Quota       uint                  `form:"quota" binding:"required"`
	EventDate   time.Time             `form:"date" binding:"required" time_format:"02-01-2006"`
	EventTime   string                `form:"time" binding:"required"`
	Location    string                `form:"location" binding:"required"`
	PhotoIMG    *multipart.FileHeader `form:"photo" binding:"required"`
	StrDate     string
}

type GetWebinars struct {
	ID          uuid.UUID `json:"id"`
	WebinarName string    `json:"name"`
	Subheader   string    `json:"subheader"`
	Price       uint64    `json:"price"`
	StrDate     string    `json:"date"`
	EventTime   string    `json:"time"`
	Location    string    `json:"location"`
	Photolink   string    `json:"photolink"`
}

type WebinarInfo struct {
	ID          uuid.UUID `json:"id"`
	WebinarName string    `json:"name"`
	Price       uint64    `json:"price"`
}
