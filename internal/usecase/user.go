package usecase

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"github.com/valentinusdelvin/velo-mom-api/utils/bcrypt"
	"github.com/valentinusdelvin/velo-mom-api/utils/jwt"
	"github.com/valentinusdelvin/velo-mom-api/utils/supabase"
)

type InterUserUsecase interface {
	Register(req models.UserRegister) error
	Login(req models.UserLogin) (models.UserLoginResponse, error)
	GetUser(param models.UserParam) (entity.User, error)
	UpdateUser(param models.UserUpdate, user entity.User) error
	UpdateProfilePhoto(param models.UpdateProfilePhoto, user entity.User) error
}

type UserUsecase struct {
	ursc    repository.InterUserRepository
	bcrypt  bcrypt.InterBcrypt
	jwtAuth jwt.InterJWT
	sb      supabase.InterSupabase
}

func NewUserUsecase(userRepo repository.InterUserRepository, bcrypt bcrypt.InterBcrypt, jwtAuth jwt.InterJWT, supabase supabase.InterSupabase) InterUserUsecase {
	return &UserUsecase{
		ursc:    userRepo,
		bcrypt:  bcrypt,
		jwtAuth: jwtAuth,
		sb:      supabase,
	}
}

func (u *UserUsecase) Register(param models.UserRegister) error {
	hashedPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}

	param.ID = uuid.New()
	param.Password = hashedPassword

	user := entity.User{
		ID:          param.ID,
		DisplayName: param.DisplayName,
		Password:    hashedPassword,
		Email:       param.Email,
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
		Email: param.Email,
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

func (u *UserUsecase) GetUser(param models.UserParam) (entity.User, error) {
	return u.ursc.GetUser(param)
}

func (u *UserUsecase) UpdateUser(param models.UserUpdate, user entity.User) error {
	err := u.ursc.UpdateUser(param, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) UpdateProfilePhoto(param models.UpdateProfilePhoto, user entity.User) error {
	ext := filepath.Ext(param.PhotoIMG.Filename)
	if ext == "" {
		return errors.New("invalid file extension: no file extension found")
	}
	param.PhotoIMG.Filename = fmt.Sprintf("%s%s", param.ID.String(), ext)

	newPhotoLink, err := u.sb.Upload(param.PhotoIMG)
	if err != nil {
		return err
	}

	if param.PhotoLink != "" {
		_ = u.sb.Delete(param.PhotoLink)
	}

	param.PhotoLink = newPhotoLink

	err = u.ursc.UpdateProfilePhoto(param, user)
	if err != nil {
		return err
	}
	return nil
}
