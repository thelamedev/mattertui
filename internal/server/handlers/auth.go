package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/thelamedev/mattertui/internal/server/store"
	"github.com/thelamedev/mattertui/pkg/shared/util"
)

const jwtExpiry = 24 * time.Hour

type registerRequestBody struct {
	Username string `json:"username" binding:"required,min=3,max=32"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}

func HandleRegisterUser(c *gin.Context) {
	var body registerRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	passwordHash, err := util.GeneratePasswordHash(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser, err := queries.CreateUser(ctx, store.CreateUserParams{
		Username:     body.Username,
		Email:        body.Email,
		PasswordHash: passwordHash,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := util.GenerateToken(jwt.MapClaims{
		"user_id": newUser.ID,
		"version": newUser.PasswordVersion,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(jwtExpiry).Unix(),
	}, cfg.Auth.JWTSecret)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": newUser, "token": token})
}

type loginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleLoginUser(c *gin.Context) {
	var body loginUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	user, err := queries.GetUserByUsername(ctx, body.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if err := util.ComparePasswordHash(user.PasswordHash, body.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token, err := util.GenerateToken(jwt.MapClaims{
		"user_id": user.ID,
		"version": user.PasswordVersion,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(jwtExpiry).Unix(),
	}, cfg.Auth.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}
