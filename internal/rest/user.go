package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
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
				"error":   err.Error(),
			})
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "user already exists",
				"error":   err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to register user",
				"error":   err.Error(),
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

func (r *Rest) GetUserInfo(ctx *gin.Context) {
	user, ok := ctx.Get("userID")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	user, err := r.usecase.UserUsecase.GetUserInfo(user.(uuid.UUID))
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
				"message": "failed to get user",
				"error":   err,
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, user)

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

	user, ok := ctx.Get("userID")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "failed to get user",
		})
		return
	}

	err = r.usecase.UserUsecase.UpdateUser(param, user.(uuid.UUID))
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

func (r *Rest) UpdateProfilePhoto(ctx *gin.Context) {
	param := models.UpdateProfilePhoto{}

	err := ctx.ShouldBindWith(&param, binding.FormMultipart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request body",
			"error":   err,
		})
		return
	}

	user, ok := ctx.Get("userID")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "failed to get user",
		})
		return
	}

	err = r.usecase.UserUsecase.UpdateProfilePhoto(param, user.(uuid.UUID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
				"error":   err.Error(),
			})
		} else if errors.Is(err, gorm.ErrInvalidData) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to update user",
				"error":   err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to update user",
				"error":   err.Error(),
			})
		}
		return
	}
	fmt.Println(user.(uuid.UUID))
	ctx.JSON(http.StatusOK, gin.H{})
}
