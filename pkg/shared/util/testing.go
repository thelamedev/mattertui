package util

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func AuthenticateTestUser(t *testing.T, c *gin.Context) {
	user_id := "56132e37-e2c7-4d1b-bcae-0ac8ecf0eeaf"

	c.Set("user_id", user_id)
	c.Set("version", 1)
}
