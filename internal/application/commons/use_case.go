package commons

import "github.com/gin-gonic/gin"

type UseCase interface {
	Execute(ctx *gin.Context)
}
