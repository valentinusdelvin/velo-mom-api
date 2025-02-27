package usecase

import (
	"github.com/valentinusdelvin/velo-mom-api/addition/bcrypt"
	"github.com/valentinusdelvin/velo-mom-api/addition/jwt"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/models"
)

type InterUserUsecase interface {
	Register(req models.UserRegister) error
	Login(req models.UserLogin) (models.UserLoginResponse, error)
}

type UserUsecase struct {
	ursc    repository.InterUserRepository
	bcrypt  bcrypt.InterBcrypt
	jwtAuth jwt.InterJWT
}

func NewUserUsecase(userRepo repository.InterUserRepository, bcrypt bcrypt.InterBcrypt, jwtAuth jwt.InterJWT) InterUserUsecase {
	return &UserUsecase{
		ursc:    userRepo,
		bcrypt:  bcrypt,
		jwtAuth: jwtAuth,
	}
}

func (u *UserUsecase) Register(param models.UserRegister) error {
	hashedPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}

	user := entity.User{
		ID:       param.ID,
		Username: param.Username,
		Password: hashedPassword,
		Email:    param.Email,
	}

	_, err = u.ursc.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) Login(param models.UserLogin) (models.UserLoginResponse, error) {
	result := models.UserLoginResponse{}

	user, err := u.ursc.GetUser(models.UserParam{
		Username: param.Username,
	})
	if err != nil {
		return result, err
	}

	err = u.bcrypt.CompareHashAndPassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := u.jwtAuth.CreateToken(user.ID)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}
