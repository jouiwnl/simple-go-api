package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jouiwnl/go-crud/initializers"
	"github.com/jouiwnl/go-crud/routes"
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
	routes.GetRoutes(r)
	r.Run()
}
