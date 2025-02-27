package jwt

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
)

type InterJWT interface {
	CreateToken(UserId uuid.UUID) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, error)
	GetLoginUser(ctx *gin.Context) (entity.User, error)
}

type JWTService struct {
	SecretKey string
	Expired   time.Duration
}

type Claims struct {
	UserId uuid.UUID
	jwt.RegisteredClaims
}

func NewJWT() InterJWT {
	secretKey := os.Getenv("JWT_SEC_KEY")
	expiredTime, err := strconv.Atoi(os.Getenv("JWT_EXP_TIME"))
	if err != nil {
		log.Panicf("Failed to set expired time for token: %v", err.Error())
	}

	return &JWTService{
		SecretKey: secretKey,
		Expired:   time.Duration(expiredTime) * time.Minute,
	}
}

func (w *JWTService) CreateToken(UserId uuid.UUID) (string, error) {
	claims := &Claims{
		UserId: UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(w.Expired)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func (w *JWTService) ValidateToken(tokenString string) (uuid.UUID, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(w.SecretKey), nil
	})

	if err != nil {
		return claims.UserId, err
	}

	if !token.Valid {
		return claims.UserId, err
	}

	return claims.UserId, nil
}

func (w *JWTService) GetLoginUser(ctx *gin.Context) (entity.User, error) {
	user, ok := ctx.Get("user")
	if !ok {
		return entity.User{}, errors.New("User not found")
	}

	return user.(entity.User), nil
}
