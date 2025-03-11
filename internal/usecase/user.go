package usecase

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"github.com/valentinusdelvin/velo-mom-api/pkg/bcrypt"
	"github.com/valentinusdelvin/velo-mom-api/pkg/jwt"
	"github.com/valentinusdelvin/velo-mom-api/pkg/supabase"
)

type InterUserUsecase interface {
	Register(req models.UserRegister) error
	Login(req models.UserLogin) (models.UserLoginResponse, error)
	GetUser(param models.UserParam) (entity.User, error)
	GetUserInfo(id uuid.UUID) (models.UserInfo, error)
	UpdateUser(param models.UserUpdate, id uuid.UUID) error
	UpdateProfilePhoto(param models.UpdateProfilePhoto, id uuid.UUID) error
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

	if strings.Split(param.Email, "@")[1] == "velomom.id" {
		user.IsAdmin = true
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

	token, err := u.jwtAuth.CreateToken(user.ID, user.IsAdmin)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (u *UserUsecase) GetUser(param models.UserParam) (entity.User, error) {
	return u.ursc.GetUser(param)
}

func (u *UserUsecase) GetUserInfo(id uuid.UUID) (models.UserInfo, error) {
	return u.ursc.GetUserInfo(id)
}

func (u *UserUsecase) UpdateUser(param models.UserUpdate, id uuid.UUID) error {
	err := u.ursc.UpdateUser(param, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) UpdateProfilePhoto(param models.UpdateProfilePhoto, id uuid.UUID) error {
	ext := filepath.Ext(param.PhotoIMG.Filename)
	if ext == "" {
		return errors.New("invalid file extension: no file extension found")
	}
	param.PhotoIMG.Filename = fmt.Sprintf("%s%v%s", id.String(), time.Now().Unix(), ext)

	newPhotoLink, err := u.sb.Upload(param.PhotoIMG)
	if err != nil {
		return err
	}

	if param.PhotoLink != "" {
		_ = u.sb.Delete(param.PhotoLink)
	}

	param.PhotoLink = newPhotoLink

	err = u.ursc.UpdateProfilePhoto(param, id)
	if err != nil {
		return err
	}
	return nil
}
