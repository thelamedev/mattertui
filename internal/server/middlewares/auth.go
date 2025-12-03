package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("authorization")
	if authorization == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	token := strings.Replace(authorization, "Bearer ", "", 1)

	c.Set("user_id", token)

	c.Next()
}
