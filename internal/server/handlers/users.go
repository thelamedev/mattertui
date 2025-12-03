package handlers

import "github.com/gin-gonic/gin"

func HandleMeProfile(c *gin.Context) {
	user_id, ok := c.Get("user_id")
	if !ok || user_id == "" {
		c.JSON(401, gin.H{
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

	c.JSON(200, gin.H{
		"user_id": id,
	})
}
