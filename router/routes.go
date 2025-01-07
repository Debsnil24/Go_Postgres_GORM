package router

import (
	"github.com/Debsnil24/Go_Postgres_GORM/middleware"
	"github.com/gofiber/fiber/v2"
)

type Repository struct {
	middleware.Repository
}

func(r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_book", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
	api.Put("/update_book/:id", r.UpdateBook)
}