package main

import (
	"log"
	"os"

	"github.com/Debsnil24/Go_Postgres_GORM/middleware"
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
		Config: models.Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Password: os.Getenv("DB_PASS"),
			User:     os.Getenv("DB_USER"),
			SSLMode:  os.Getenv("DB_SSL"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Couldn't Load the Database")
	}

	err = middleware.MigrateBooks(db)
	if err != nil {
		log.Fatal("Could not migrate DB")
	}

	router := router.Repository{
		Repository: middleware.Repository{
			Repository: models.Repository{
				DB: db,
			},
		},
	}

	app := fiber.New()
	router.SetupRoutes(app)
	log.Println("Starting Server on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
