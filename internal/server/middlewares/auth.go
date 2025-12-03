package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/thelamedev/mattertui/internal/config"
	"github.com/thelamedev/mattertui/internal/server/database"
	"github.com/thelamedev/mattertui/internal/server/store"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	pool := database.GetPool()
	queries := store.New(pool)

	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header is missing or invalid",
			})
			return
		}

		token := strings.Replace(authorization, "Bearer ", "", 1)

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
			return []byte(cfg.Auth.JWTSecret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "failed to parse token",
			})
			return
		}

		var userID pgtype.UUID
		if err := userID.Scan(claims["user_id"]); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "failed to parse user_id from token",
			})
			return
		}

		pwdVersion, err := queries.GetUserPasswordVersionByUserID(c.Request.Context(), userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "failed to get user password version",
			})
			return
		}

		log.Printf("claims: %v", claims)
		log.Printf("pwdVersion: %v", pwdVersion)

		claimedVersion, ok := claims["version"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "version claim is missing",
			})
			return
		}

		claimedVersionInt, ok := claimedVersion.(int32)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "version claim is invalid",
			})
			return
		}

		if claimedVersionInt != pwdVersion {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "password version mismatch",
			})
			return
		}

		c.Set("user_id", userID)
		c.Set("version", claimedVersionInt)

		c.Next()
	}
}
