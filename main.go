package main

import (
	"github.com/valentinusdelvin/velo-mom-api/initializers"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/internal/rest"
	"github.com/valentinusdelvin/velo-mom-api/internal/usecase"
	"github.com/valentinusdelvin/velo-mom-api/pkg/bcrypt"
	"github.com/valentinusdelvin/velo-mom-api/pkg/jwt"
	"github.com/valentinusdelvin/velo-mom-api/pkg/middleware"
	"github.com/valentinusdelvin/velo-mom-api/pkg/supabase"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	bcrypt := bcrypt.Init()
	jwt := jwt.NewJWT()
	db := initializers.ConnectToDB()
	supabase := supabase.Init()
	initializers.AutoMigrate(db)

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(usecase.InitializersParam{
		Repository: repository,
		Bcrypt:     &bcrypt,
		JWT:        &jwt,
		Supabase:   &supabase,
	})
	middleware := middleware.Init(usecase)

	rest := rest.NewRest(usecase, middleware)
	rest.FinalCheck()

	rest.Run()
}
