package usecase

import (
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/pkg/bcrypt"
	"github.com/valentinusdelvin/velo-mom-api/pkg/jwt"
	"github.com/valentinusdelvin/velo-mom-api/pkg/supabase"
)

type Usecase struct {
	UserUsecase    InterUserUsecase
	ArticleUsecase InterArticleUsecase
	VideoUsecase   InterVideoUsecase
	JournalUsecase InterJournalUsecase
	WebinarUsecase InterWebinarUsecase
}

type InitializersParam struct {
	Repository *repository.Repository
	Bcrypt     *bcrypt.InterBcrypt
	JWT        *jwt.InterJWT
	Supabase   *supabase.InterSupabase
}

func NewUsecase(param InitializersParam) *Usecase {
	UserUsecase := NewUserUsecase(param.Repository.UserRepository, *param.Bcrypt, *param.JWT, *param.Supabase)
	ArticleUsecase := NewArticleUsecase(param.Repository.ArticleRepository, *param.Supabase)
	VideoUsecase := NewVideoUsecase(param.Repository.VideoRepository)
	JournalUsecase := NewJournalUsecase(param.Repository.JournalRepository)
	WebinarUsecase := NewWebinarUsecase(param.Repository.WebinarRepository, *param.Supabase)

	return &Usecase{
		UserUsecase:    UserUsecase,
		ArticleUsecase: ArticleUsecase,
		VideoUsecase:   VideoUsecase,
		JournalUsecase: JournalUsecase,
		WebinarUsecase: WebinarUsecase,
	}
}
