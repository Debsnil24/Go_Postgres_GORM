package main

import (
	"log"

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

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Couldn't Load the Database")
	}

	r := router.Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}