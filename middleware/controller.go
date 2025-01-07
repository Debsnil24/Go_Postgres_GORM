package middleware

import (
	"net/http"

	"github.com/Debsnil24/Go_Postgres_GORM/models"
	"github.com/gofiber/fiber/v2"
)
type Repository struct {
	models.Repository
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

func(r *Repository) GetBooks(context *fiber.Ctx) error {
	book := &[]models.Books{}
	err := r.DB.Find(book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Unable to Find Book"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Book found Successfully",
			"data" : book,
	})
	return nil
}

func(r *Repository) GetBookByID(context *fiber.Ctx) error {
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Book Found",
			"data":  " ",
		})
	return nil
}

func(r *Repository) DeleteBook(context *fiber.Ctx) error {
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Book Deleted Successfully",
			"data":  " ",
		})
	return nil
}