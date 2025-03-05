package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/models"
	addition "github.com/valentinusdelvin/velo-mom-api/utils/timeconvert"
)

type InterJournalUsecase interface {
	CreateJournal(param models.CreateJournal) error
	GetUserJournals(userID uuid.UUID) ([]models.CreateJournal, error)
	GetUserJournalByID(userID uuid.UUID, id string) (entity.Journal, error)
}

type JournalUsecase struct {
	jrsc repository.InterJournalRepository
}

func NewJournalUsecase(journalRepo repository.InterJournalRepository) InterJournalUsecase {
	return &JournalUsecase{
		jrsc: journalRepo,
	}
}

func (j *JournalUsecase) CreateJournal(param models.CreateJournal) error {
	journalpost := entity.Journal{
		ID:        uuid.New(),
		UserID:    param.UserID,
		Title:     param.Title,
		Story:     param.Story,
		Feels:     param.Feels,
		Emoji:     entity.Emoji(param.Emoji),
		CreatedAt: addition.TimeConvert(time.Now()),
	}

	_, err := j.jrsc.CreateJournal(journalpost)
	if err != nil {
		return err
	}
	return nil
}

func (j *JournalUsecase) GetUserJournals(userID uuid.UUID) ([]models.CreateJournal, error) {
	journals, err := j.jrsc.GetUserJournals(userID)
	if err != nil {
		return nil, err
	}
	return journals, nil
}

func (j *JournalUsecase) GetUserJournalByID(userID uuid.UUID, id string) (entity.Journal, error) {
	journal, err := j.jrsc.GetUserJournalByID(userID, id)
	if err != nil {
		return journal, err
	}
	return journal, nil
}
