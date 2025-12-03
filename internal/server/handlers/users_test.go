package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUsersHandler(t *testing.T) {
	t.Run("authenticated requests", func(t2 *testing.T) {
		t2.Parallel()

		t2.Run("get own profile", func(t3 *testing.T) {
			t3.Parallel()
			req := httptest.NewRequestWithContext(context.Background(), "GET", "/api/v1/users/me", nil)

			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)
			c.Request = req
			c.Set("user_id", "test")
			HandleMeProfile(c)

			assert.Equal(t3, http.StatusOK, recorder.Code)
		})

		t2.Run("get public profile", func(t3 *testing.T) {
			t3.Parallel()
			req := httptest.NewRequestWithContext(context.Background(), "GET", "/api/v1/users/1", nil)

			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)
			c.Request = req
			c.Set("user_id", "test")
			HandleGetUserByID(c)

			assert.Equal(t3, http.StatusOK, recorder.Code)
		})
	})

	t.Run("unauthenticated requests", func(t2 *testing.T) {
		t2.Parallel()

		t2.Run("get own profile", func(t3 *testing.T) {
			t3.Parallel()
			req := httptest.NewRequestWithContext(context.Background(), "GET", "/api/v1/users/me", nil)

			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)
			c.Request = req
			HandleMeProfile(c)

			assert.Equal(t3, http.StatusUnauthorized, recorder.Code)
		})

		t2.Run("get public profile", func(t3 *testing.T) {
			t3.Parallel()
			req := httptest.NewRequestWithContext(context.Background(), "GET", "/api/v1/users/1", nil)

			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)
			c.Request = req
			HandleGetUserByID(c)

			assert.Equal(t3, http.StatusOK, recorder.Code)
		})
	})
}
