package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jouiwnl/go-crud/controllers"
	"github.com/jouiwnl/go-crud/initializers"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Ocorreu um erro ao tentar iniciar a aplicação")
	}

	initializers.ConnectDb()
}

func main() {

	r := gin.Default()
	r.GET("/users/:id", controllers.GetUser)
	r.GET("/users", controllers.GetUSers)

	r.Run() // listen and serve on 0.0.0.0:8080
}
