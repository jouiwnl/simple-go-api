package main

import (
	"database/sql"
	"github.com/jouiwnl/simple-go-api/internal/config/server"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	webServer := server.SetupRouter(dbcon)

	err = webServer.Run()

	if err != nil {
		panic(err.Error())
	}
}
