package repository

import (
	"fmt"

	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"gorm.io/gorm"
)

type InterVideoRepository interface {
	CreateVideo(video entity.Video) (entity.Video, error)
	GetVideos(page, size int) ([]models.CreateVideo, error)
	GetVideoByID(id string) (entity.Video, error)
	GetVideosBySearch(param models.CreateVideo, page, size int) ([]entity.Video, error)
	GetVideoByFilter(param models.CreateVideo, page, size int) ([]entity.Video, error)
}

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) InterVideoRepository {
	return &VideoRepository{
		db: db,
	}
}

func (vr *VideoRepository) CreateVideo(video entity.Video) (entity.Video, error) {
	err := vr.db.Create(&video).Error
	if err != nil {
		return video, err
	}
	return video, nil
}

func (vr *VideoRepository) GetVideos(page, size int) ([]models.CreateVideo, error) {
	var videos []models.CreateVideo
	offset := (page - 1) * size

	err := vr.db.Model(entity.Video{}).Order("created_at DESC").Limit(size).Offset(offset).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (vr *VideoRepository) GetVideoByID(id string) (entity.Video, error) {
	var video entity.Video

	err := vr.db.Model(entity.Video{}).Where("youtube_id = ?", id).Find(&video).Error
	if err != nil {
		return video, err
	}

	return video, nil
}

func (vr *VideoRepository) GetVideosBySearch(param models.CreateVideo, page, size int) ([]entity.Video, error) {
	var videos []entity.Video

	offset := (page - 1) * size

	err := vr.db.Model(entity.Video{}).Order("created_at DESC").Limit(size).Offset(offset).Where("title ILIKE ?", fmt.Sprintf("%%%s%%", param.Title)).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (vr *VideoRepository) GetVideoByFilter(param models.CreateVideo, page, size int) ([]entity.Video, error) {
	var videos []entity.Video

	offset := (page - 1) * size

	err := vr.db.Model(entity.Video{}).Order("created_at DESC").Limit(size).Offset(offset).Where("filter = ?", param.Filter).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
