package usecase

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"github.com/valentinusdelvin/velo-mom-api/pkg/supabase"
	"github.com/valentinusdelvin/velo-mom-api/pkg/timeconvert"
)

type InterWebinarUsecase interface {
	CreateWebinar(param models.CreateWebinar) error
	DeleteWebinar(id string) error
	GetWebinars(page, size int) ([]models.GetWebinars, error)
	GetWebinarByID(id string) (entity.Webinar, error)
	GetPurchasedWebinars(userID uuid.UUID) ([]entity.Webinar, error)
}

type WebinarUsecase struct {
	wrsc repository.InterWebinarRepository
	sb   supabase.InterSupabase
}

func NewWebinarUsecase(WebinarRepo repository.InterWebinarRepository, supabase supabase.InterSupabase) InterWebinarUsecase {
	return &WebinarUsecase{
		wrsc: WebinarRepo,
		sb:   supabase,
	}
}

func (w *WebinarUsecase) CreateWebinar(param models.CreateWebinar) error {
	param.ID = uuid.New()
	ext := filepath.Ext(param.PhotoIMG.Filename)
	if ext == "" {
		return errors.New("invalid file extension: no file extension found")
	}
	param.PhotoIMG.Filename = fmt.Sprintf("%s-%v%s", param.ID, time.Now().Unix(), ext)

	newPhotoLink, err := w.sb.Upload(param.PhotoIMG)
	if err != nil {
		fmt.Println("gagal upload")
		return errors.New("can't upload")
	}

	webinarPost := entity.Webinar{
		ID:          param.ID,
		WebinarName: param.WebinarName,
		Subheader:   param.Subheader,
		Description: param.Description,
		Price:       param.Price,
		Photolink:   newPhotoLink,
		Quota:       param.Quota,
		StrDate:     timeconvert.TimeConvert(param.EventDate),
		EventTime:   param.EventTime,
		Location:    param.Location,
		CreatedAt:   time.Now(),
	}

	_, err = w.wrsc.CreateWebinar(webinarPost)
	if err != nil {
		return errors.New("failed to create webinar: " + err.Error())
	}
	return nil
}

func (w *WebinarUsecase) DeleteWebinar(id string) error {
	webinar, err := w.wrsc.GetWebinarByID(id)
	if err != nil {
		return err
	}

	err = w.sb.Delete(webinar.Photolink)
	if err != nil {
		return err
	}

	return w.wrsc.DeleteWebinar(id)
}

func (w *WebinarUsecase) GetWebinars(page, size int) ([]models.GetWebinars, error) {
	webinars, err := w.wrsc.GetWebinars(page, size)
	if err != nil {
		return nil, err
	}
	return webinars, nil
}

func (w *WebinarUsecase) GetWebinarByID(id string) (entity.Webinar, error) {
	webinar, err := w.wrsc.GetWebinarByID(id)
	if err != nil {
		return entity.Webinar{}, err
	}
	return webinar, nil
}

func (w *WebinarUsecase) GetPurchasedWebinars(userID uuid.UUID) ([]entity.Webinar, error) {
	return w.wrsc.GetPurchasedWebinars(userID)
}
