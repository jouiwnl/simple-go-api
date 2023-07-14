package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUtils(t *testing.T) {
	t.Run("should write an internal error response", func(t *testing.T) {
		err := "Internal Server Error"

		fakeServer := gin.Default()
		fakeServer.GET("/test", func(ctx *gin.Context) {
			WriteInternalErrorResponse(ctx, err)
		})

		request, _ := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		fakeServer.ServeHTTP(w, request)

		assert.Equal(t, 500, w.Code)
	})

	t.Run("should write an not found error response", func(t *testing.T) {
		fakeServer := gin.Default()
		fakeServer.GET("/test", WriteNotFoundErrorResponse)

		request, _ := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		fakeServer.ServeHTTP(w, request)

		assert.Equal(t, 404, w.Code)
	})

	t.Run("should return empty string when value is null", func(t *testing.T) {
		nullSafe := NullSafeString(nil)

		assert.Equal(t, nullSafe, "")
	})

	t.Run("should return the same string when value is a valid string", func(t *testing.T) {
		anyString := "test"
		nullSafe := NullSafeString(&anyString)

		assert.Equal(t, nullSafe, "test")
	})
}
