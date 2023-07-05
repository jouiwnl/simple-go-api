package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func init() {
	godotenv.Load()
}

func ConnectDb() {
	configuration := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_SECRET"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PORT"),
	)

	var err error
	Database, err = gorm.Open(postgres.Open(configuration), &gorm.Config{})

	if err != nil {
		log.Fatal("Falha ao conectar-se com o banco de dados")
	}
}
