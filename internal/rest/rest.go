package rest

import (
	"log"
	"net/http"

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
	routerGroup.GET("/login-user", r.middleware.Authenticate, getLoginUser)
	routerGroup.POST("/auth-email", r.middleware.Authenticate, r.AuthenticateEmail)
	routerGroup.PATCH("/update-user", r.middleware.Authenticate, r.UpdateUser)
	routerGroup.PATCH("/update-photo", r.middleware.Authenticate, r.UpdateProfilePhoto)

	articlepost := routerGroup.Group("/articles")
	articlepost.POST("/", r.middleware.Authenticate, r.CreateArticle)
	articlepost.GET("/", r.GetArticles)
	articlepost.GET("/:id", r.GetArticleByID)

	videopost := routerGroup.Group("/videos")
	videopost.POST("/", r.middleware.Authenticate, r.CreateVideo)
	videopost.GET("/", r.GetVideos)
	videopost.GET("/:id", r.GetVideoByID)

	journalpost := routerGroup.Group("/journals")
	journalpost.POST("/", r.middleware.Authenticate, r.CreateJournal)
	journalpost.GET("/me", r.middleware.Authenticate, r.GetUserJournals)
	journalpost.GET("/me/:id", r.middleware.Authenticate, r.GetUserJournalByID)
}

func (r *Rest) Run() {
	r.router.Run(":8080")

	err := r.router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func getLoginUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "failed to login"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
