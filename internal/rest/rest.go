package rest

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/velo-mom-api/internal/usecase"
	"github.com/valentinusdelvin/velo-mom-api/pkg/middleware"
)

type Rest struct {
	router     *gin.Engine
	usecase    *usecase.Usecase
	middleware middleware.Interface
}

func NewRest(usecase *usecase.Usecase, middleware middleware.Interface) *Rest {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	return &Rest{
		router:     gin.Default(),
		usecase:    usecase,
		middleware: middleware,
	}
}

func (r *Rest) FinalCheck() {
	routerGroup := r.router.Group("/api")

	routerGroup.POST("/register", r.Register)
	routerGroup.POST("/login", r.Login)
	routerGroup.GET("/me", r.middleware.Authenticate, r.GetUserInfo)
	routerGroup.POST("/auth-email", r.middleware.Authenticate, r.AuthenticateEmail)
	routerGroup.PATCH("/update-user", r.middleware.Authenticate, r.UpdateUser)
	routerGroup.PATCH("/update-photo", r.middleware.Authenticate, r.UpdateProfilePhoto)

	articlepost := routerGroup.Group("/articles")
	articlepost.POST("/", r.middleware.Authenticate, r.middleware.Authorization, r.CreateArticle)
	articlepost.GET("/", r.GetArticles)
	articlepost.GET("/:id", r.GetArticleByID)
	articlepost.GET("/search", r.GetArticlesBySearch)

	videopost := routerGroup.Group("/videos")
	videopost.POST("/", r.middleware.Authenticate, r.middleware.Authorization, r.CreateVideo)
	videopost.GET("/", r.GetVideos)
	videopost.GET("/:id", r.GetVideoByID)
	videopost.GET("/search", r.GetVideosBySearch)

	journalpost := routerGroup.Group("/journals")
	journalpost.POST("/", r.middleware.Authenticate, r.CreateJournal)
	journalpost.GET("/me", r.middleware.Authenticate, r.GetUserJournals)
	journalpost.GET("/me/:id", r.middleware.Authenticate, r.GetUserJournalByID)

	webinarpost := routerGroup.Group("/webinars")
	webinarpost.POST("/", r.middleware.Authenticate, r.middleware.Authorization, r.CreateWebinar)
	webinarpost.GET("/", r.GetWebinars)
	webinarpost.GET("/:id", r.GetWebinarByID)
	webinarpost.GET("/my-webinars", r.middleware.Authenticate, r.GetPurchasedWebinars)
	webinarpost.POST("/purchase/:id", r.middleware.Authenticate, r.Purchase)
	webinarpost.POST("/validate", r.Validate)
}

func (r *Rest) Run() {

	err := r.router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
