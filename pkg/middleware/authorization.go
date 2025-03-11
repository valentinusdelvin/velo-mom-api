package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Authorization(c *gin.Context) {
	IsAdmin, exists := c.Get("isAdmin")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can't find the user's role",
		})
		c.Abort()
		return
	}

	admin, ok := IsAdmin.(bool)
	if !ok || !admin {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "forbidden",
		})
		c.Abort()
		return
	}

	c.Next()
}
