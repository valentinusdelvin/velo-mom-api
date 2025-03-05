package models

type CreateVideo struct {
	YoutubeID   string `json:"id"`
	Title       string `json:"title" binding:"required"`
	YoutubeURL  string `json:"videoURL" binding:"required"`
	Description string `json:"description" binding:"required"`
	Thumbnail   string `json:"thumbnail"`
}
