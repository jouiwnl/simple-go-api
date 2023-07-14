package utils

import "github.com/gin-gonic/gin"

type Utils struct {
}

func WriteInternalErrorResponse(ginCtx *gin.Context, error string) {
	ginCtx.JSON(500, gin.H{
		"message": "Internal Server Error",
		"error":   error,
	})
}

func WriteNotFoundErrorResponse(ginCtx *gin.Context) {
	ginCtx.JSON(404, gin.H{
		"message": "Not found",
		"error":   "Not found",
	})
}

func NullSafeString(original *string) string {
	if original == nil {
		return ""
	}

	return *original
}
