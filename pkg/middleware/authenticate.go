package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Authenticate(c *gin.Context) {
	bearertoken := c.GetHeader("Authorization")
	if bearertoken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "token is required",
		})
		c.Abort()
		return
	}

	token := strings.Split(bearertoken, " ")[1]
	userId, err := m.jwt.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token/Failed to validate token",
		})
		c.Abort()
		return
	}

	c.Set("userID", userId)

	c.Next()
}
