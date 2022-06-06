package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pharaohthecat/blog-application-go/blogbackend/database"
	"github.com/pharaohthecat/blog-application-go/blogbackend/routes"
)

func main(){
	database.Connect()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":"+port)
}