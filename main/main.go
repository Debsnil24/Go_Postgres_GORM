package main

import (
	"log"
	"os"

	"github.com/Debsnil24/Go_Postgres_GORM/models"
	"github.com/Debsnil24/Go_Postgres_GORM/router"
	"github.com/Debsnil24/Go_Postgres_GORM/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load .env file")
	}

	config := &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		SSLMode: os.Getenv("DB_SSL"),
		DBName: os.Getenv("DB_NAME"),
	}
	
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Couldn't Load the Database")
	}

	r := models.Repository{
		DB: db,
	}
	
	router := *&router.Repository{
		Repository: r,
	}

	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(":8080")
}