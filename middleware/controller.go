package middleware

import (
	"net/http"

	"github.com/Debsnil24/Go_Postgres_GORM/models"
	"github.com/Debsnil24/Go_Postgres_GORM/router"
	"github.com/gofiber/fiber/v2"
)
type Repository struct {
	router.Repository
}

func(r *Repository) CreateBook(context *fiber.Ctx) error {
	book := models.Book{}
	err := context.BodyParser(&book)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request Failed"})
		return err
	}
	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Unable to Create Book"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Book Created Successfully"})
		return nil
}