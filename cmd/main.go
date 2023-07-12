package main

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jouiwnl/simple-go-api/internal/config/database/sqlc"
	"github.com/jouiwnl/simple-go-api/internal/factory"
	userUseCases "github.com/jouiwnl/simple-go-api/internal/user/usecases"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Ocorreu um erro ao tentar iniciar a aplicação")
	}
}

func main() {
	dbcon, err := sql.Open(
		"postgres",
		os.Getenv("POSTGRES_DATABASE_URL"),
	)

	if err != nil {
		panic(err.Error())
	}

	queries := *sqlc.New(dbcon)
	ctx := context.Background()

	repositoryFactory := factory.NewRepositoryFactory(dbcon, queries, ctx)

	userRepository := repositoryFactory.NewUserRepository()

	getAllUsersUseCase := userUseCases.NewGetAllUsersUseCase(userRepository)
	getUserByIdUseCase := userUseCases.NewGetUserByIdUseCase(userRepository)
	createUserUseCase := userUseCases.NewCreateUserUseCase(userRepository)
	updateUserUseCase := userUseCases.NewUpdateUserUseCase(userRepository)

	server := gin.Default()

	server.GET("/users/:id", getUserByIdUseCase.GetUserById)
	server.GET("/users", getAllUsersUseCase.GetPaginatedUsers)
	server.POST("/users", createUserUseCase.CreateUser)
	server.PUT("/users/:id", updateUserUseCase.UpdateUser)

	server.Run()
}
