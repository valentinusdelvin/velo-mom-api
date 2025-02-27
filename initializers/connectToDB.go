package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("POSTGRESQL_DSN"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect to DB")
	}
	fmt.Println("Connected to DB")
	return db
}
