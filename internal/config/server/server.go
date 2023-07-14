package server

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jouiwnl/simple-go-api/internal/config/database/sqlc"
	"github.com/jouiwnl/simple-go-api/internal/factory"
	userUseCases "github.com/jouiwnl/simple-go-api/internal/user/usecases"
)

func SetupRouter(dbcon *sql.DB) *gin.Engine {
	r := gin.Default()

	queries := *sqlc.New(dbcon)
	ctx := context.Background()

	repositoryFactory := factory.NewRepositoryFactory(dbcon, queries, ctx)

	userRepository := repositoryFactory.NewUserRepository()

	getAllUsersUseCase := userUseCases.NewGetAllUsersUseCase(userRepository)
	getUserByIdUseCase := userUseCases.NewGetUserByIdUseCase(userRepository)
	createUserUseCase := userUseCases.NewCreateUserUseCase(userRepository)
	updateUserUseCase := userUseCases.NewUpdateUserUseCase(userRepository)
	deleteUserCase := userUseCases.NewDeleteUserCase(userRepository)

	server := gin.Default()

	server.GET("/users/:id", getUserByIdUseCase.Execute)
	server.GET("/users", getAllUsersUseCase.Execute)
	server.POST("/users", createUserUseCase.Execute)
	server.PUT("/users/:id", updateUserUseCase.Execute)
	server.DELETE("/users/:id", deleteUserCase.Execute)

	return r
}
