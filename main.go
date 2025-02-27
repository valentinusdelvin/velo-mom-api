package main

import (
	"fmt"

	"github.com/valentinusdelvin/velo-mom-api/addition/bcrypt"
	"github.com/valentinusdelvin/velo-mom-api/addition/jwt"
	"github.com/valentinusdelvin/velo-mom-api/initializers"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/internal/usecase"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	bcrypt := bcrypt.Init()
	jwt := jwt.NewJWT()
	db := initializers.ConnectToDB()
	initializers.AutoMigrate(db)

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(usecase.InitializersParam{
		Repository: repository,
		Bcrypt:     &bcrypt,
		JWT:        &jwt,
	})
	fmt.Println(usecase)
}
