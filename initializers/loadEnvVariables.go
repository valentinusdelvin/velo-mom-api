package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	fmt.Println("Running LoadEnvVariables, STAGE =", os.Getenv("STAGE"))
	if os.Getenv("STAGE") == "PROD" {
		return
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
