package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/pharaohthecat/blog-application-go/blogbackend/models"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error ao carregar o arquivo .env")
	}

	// Conexão do Banco baseado no arquivo .env na raiz (DSN)
	// https://pkg.go.dev/gorm.io/gorm
	// https://gorm.io/docs/
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = database

	if err != nil {
		panic("Não foi possível conectar ao banco de dados!")
	} else {
		log.Println("Conectado ao banco de dados!")
	}

	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)

}
