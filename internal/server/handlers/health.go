package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var startupTime = time.Now()

func HandleHealthCheck(c *gin.Context) {
	uptime := time.Since(startupTime)

	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"uptime": uptime.String(),
	})
}
