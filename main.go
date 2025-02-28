package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/velo-mom-api/addition/bcrypt"
	"github.com/valentinusdelvin/velo-mom-api/addition/jwt"
	"github.com/valentinusdelvin/velo-mom-api/addition/middleware"
	"github.com/valentinusdelvin/velo-mom-api/initializers"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/internal/rest"
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
	middleware := middleware.Init(usecase)

	rest := rest.NewRest(usecase, middleware)
	rest.FinalCheck()

	r := gin.Default()
	for _, route := range r.Routes() {
		fmt.Println(route.Method, route.Path)
	}
	rest.Run()
}
