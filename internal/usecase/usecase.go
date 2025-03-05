package usecase

import (
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/utils/bcrypt"
	"github.com/valentinusdelvin/velo-mom-api/utils/jwt"
)

type Usecase struct {
	UserUsecase    InterUserUsecase
	ArticleUsecase InterArticleUsecase
	VideoUsecase   InterVideoUsecase
	JournalUsecase InterJournalUsecase
}

type InitializersParam struct {
	Repository *repository.Repository
	Bcrypt     *bcrypt.InterBcrypt
	JWT        *jwt.InterJWT
}

func NewUsecase(param InitializersParam) *Usecase {
	UserUsecase := NewUserUsecase(param.Repository.UserRepository, *param.Bcrypt, *param.JWT)
	ArticleUsecase := NewArticleUsecase(param.Repository.ArticleRepository)
	VideoUsecase := NewVideoUsecase(param.Repository.VideoRepository)
	JournalUsecase := NewJournalUsecase(param.Repository.JournalRepository)

	return &Usecase{
		UserUsecase:    UserUsecase,
		ArticleUsecase: ArticleUsecase,
		VideoUsecase:   VideoUsecase,
		JournalUsecase: JournalUsecase,
	}
}
