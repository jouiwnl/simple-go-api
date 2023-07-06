package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jouiwnl/go-crud/usecases"
)

func GetRoutes(r *gin.Engine) {
	r.GET("/users/:id", usecases.GetUserById)
	r.GET("/users", usecases.GetAllUsers)
	r.POST("/users", usecases.CreateUser)
	r.PUT("/users/:id", usecases.UpdateUser)
}
