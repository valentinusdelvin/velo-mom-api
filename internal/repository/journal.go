package repository

import (
	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"gorm.io/gorm"
)

type InterJournalRepository interface {
	CreateJournal(journal entity.Journal) (entity.Journal, error)
	GetUserJournals(userID uuid.UUID) ([]models.CreateJournal, error)
	GetUserJournalByID(userID uuid.UUID, id string) (entity.Journal, error)
}

type JournalRepository struct {
	db *gorm.DB
}

func NewJournalRepository(db *gorm.DB) InterJournalRepository {
	return &JournalRepository{
		db: db,
	}
}

func (jr *JournalRepository) CreateJournal(journal entity.Journal) (entity.Journal, error) {
	err := jr.db.Create(&journal).Error
	if err != nil {
		return journal, err
	}
	return journal, nil
}

func (jr *JournalRepository) GetUserJournals(userID uuid.UUID) ([]models.CreateJournal, error) {
	var journals []models.CreateJournal

	err := jr.db.Table("journals").Where("user_id = ?", userID).Find(&journals).Error
	if err != nil {
		return nil, err
	}

	return journals, nil
}

func (jr *JournalRepository) GetUserJournalByID(userID uuid.UUID, id string) (entity.Journal, error) {
	var journal entity.Journal

	err := jr.db.Table("journals").Where("id = ?", id).Where("user_id = ?", userID).Find(&journal).Error
	if err != nil {
		return journal, err
	}
	return journal, nil
}
