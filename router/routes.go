package router

import (
	"github.com/Debsnil24/Go_Postgres_GORM/middleware"
	"github.com/Debsnil24/Go_Postgres_GORM/models"
	"github.com/gofiber/fiber/v2"
)

type Repository struct {
	models.Repository
}

func(r *Repository) SetupRoutes(app *fiber.App) {
	middleware := *&middleware.Repository{}
	api := app.Group("/api")
	api.Post("/create_book", middleware.CreateBook)
	api.Delete("/delete_book/:id", middleware.DeleteBook)
	api.Get("/get_books/:id", middleware.GetBookByID)
	api.Get("/books", middleware.GetBooks)
}