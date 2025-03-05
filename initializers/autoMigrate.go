package initializers

import (
	"log"

	"github.com/valentinusdelvin/velo-mom-api/entity"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Article{},
		&entity.Video{},
		&entity.Journal{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate DB: %v", err)
	}
}
