package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"gorm.io/gorm"
)

type InterWebinarRepository interface {
	CreateWebinar(Webinar entity.Webinar) (entity.Webinar, error)
	DeleteWebinar(id string) error
	GetWebinars(page, size int) ([]models.GetWebinars, error)
	GetWebinarByID(id string) (entity.Webinar, error)
	CreateWebinarAttendee(tx *gorm.DB, attendee entity.WebinarAttendee) error
	UpdateWebinarInfo(tx *gorm.DB, webinarID uuid.UUID) error
	GetPurchasedWebinars(userID uuid.UUID) ([]entity.Webinar, error)
}

type WebinarRepository struct {
	db *gorm.DB
}

func NewWebinarRepository(db *gorm.DB) InterWebinarRepository {
	return &WebinarRepository{
		db: db,
	}
}

func (wr *WebinarRepository) CreateWebinar(webinar entity.Webinar) (entity.Webinar, error) {
	err := wr.db.Create(&webinar).Error
	if err != nil {
		return webinar, err
	}
	return webinar, nil
}

func (wr *WebinarRepository) DeleteWebinar(id string) error {
	return wr.db.Where("id = ?", id).Delete(&entity.Webinar{}).Error
}

func (wr *WebinarRepository) GetWebinars(page, size int) ([]models.GetWebinars, error) {
	var webinars []models.GetWebinars
	offset := (page - 1) * size

	err := wr.db.Model(entity.Webinar{}).Order("created_at DESC").Limit(size).Offset(offset).Find(&webinars).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(webinars)
	return webinars, nil
}

func (wr *WebinarRepository) GetWebinarByID(id string) (entity.Webinar, error) {
	var webinar entity.Webinar

	err := wr.db.Model(entity.Webinar{}).Omit("WebinarAttendees").Where("id = ?", id).Find(&webinar).Error
	if err != nil {
		return webinar, err
	}

	return webinar, nil
}

func (wr *WebinarRepository) CreateWebinarAttendee(tx *gorm.DB, attendee entity.WebinarAttendee) error {
	return tx.Create(&attendee).Error
}

func (wr *WebinarRepository) UpdateWebinarInfo(tx *gorm.DB, webinarID uuid.UUID) error {
	return tx.Model(&entity.Webinar{}).
		Where("id = ? AND quota > 0", webinarID).
		Updates(map[string]interface{}{
			"quota": gorm.Expr("quota - 1"),
		}).Error
}

func (wr *WebinarRepository) GetPurchasedWebinars(userID uuid.UUID) ([]entity.Webinar, error) {
	var webinars []entity.Webinar

	err := wr.db.
		Joins("JOIN webinar_attendees ON webinar_attendees.webinar_id = webinars.id").
		Where("webinar_attendees.user_id = ?", userID).
		Find(&webinars).Error

	if err != nil {
		return nil, err
	}

	return webinars, nil
}
