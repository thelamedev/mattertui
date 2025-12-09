package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/thelamedev/mattertui/pkg/shared/util"
)

func TestUsersHandler(t *testing.T) {
	t.Run("authenticated requests", func(t2 *testing.T) {
		t2.Parallel()

		t2.Run("get public profile", func(t3 *testing.T) {
			t3.Parallel()
			req := httptest.NewRequestWithContext(context.Background(), "GET", "/api/v1/users/56132e37-e2c7-4d1b-bcae-0ac8ecf0eeaf", nil)

			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)
			util.AuthenticateTestUser(t, c)
			c.Request = req
			c.AddParam("id", "56132e37-e2c7-4d1b-bcae-0ac8ecf0eeaf")
			HandleGetUserByID(c)

			assert.Equal(t3, http.StatusOK, recorder.Code, recorder.Body.String())
		})
	})

	t.Run("unauthenticated requests", func(t2 *testing.T) {
		t2.Parallel()

		t2.Run("get public profile", func(t3 *testing.T) {
			t3.Parallel()
			req := httptest.NewRequestWithContext(context.Background(), "GET", "/api/v1/users/56132e37-e2c7-4d1b-bcae-0ac8ecf0eeaf", nil)

			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)
			c.Request = req
			c.AddParam("id", "56132e37-e2c7-4d1b-bcae-0ac8ecf0eeaf")
			HandleGetUserByID(c)

			assert.Equal(t3, http.StatusOK, recorder.Code, recorder.Body.String())
		})
	})
}
