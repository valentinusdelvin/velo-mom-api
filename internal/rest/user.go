package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"gorm.io/gorm"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := models.UserRegister{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind with JSON",
			"error":   err,
		})
		return
	}

	err = r.usecase.UserUsecase.Register(param)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "data ivalid",
				"error":   err,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to register user",
				"error":   err,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (r *Rest) Login(ctx *gin.Context) {
	param := models.UserLogin{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind with JSON",
			"error":   err,
		})
		return
	}

	token, err := r.usecase.UserUsecase.Login(param)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "user not found",
				"error":   err,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to login",
				"error":   err,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (r *Rest) AuthenticateEmail(ctx *gin.Context) {
	param := models.EmailAuthenticator{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind with JSON",
			"error":   err,
		})
		return
	}

	_, err = r.usecase.UserUsecase.GetUser(models.UserParam{Email: param.Email})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
				"error":   err,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to authenticate email",
				"error":   err,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (r *Rest) UpdateUser(ctx *gin.Context) {
	param := models.UserUpdate{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind with JSON",
			"error":   err,
		})
		return
	}

	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "failed to get user",
		})
		return
	}

	err = r.usecase.UserUsecase.UpdateUser(param, user.(entity.User))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
				"error":   err,
			})
		} else if errors.Is(err, gorm.ErrInvalidData) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to update user",
				"error":   err,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to update user",
				"error":   err,
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
