package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func HandleMeProfile(c *gin.Context) {
	user_id, ok := c.Get("user_id")
	if !ok || user_id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	c.JSON(200, gin.H{
		"user_id": user_id,
	})
}

func HandleGetUserByID(c *gin.Context) {
	id := c.Param("id")

	var userID pgtype.UUID
	if err := userID.Scan(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	user, err := queries.GetUserByID(ctx, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
