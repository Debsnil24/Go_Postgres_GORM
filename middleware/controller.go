package middleware

import (
	"fmt"
	"net/http"

	"github.com/Debsnil24/Go_Postgres_GORM/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)
type Repository struct {
	models.Repository
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Books{})
	return err
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
	book := &models.Books{}

	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}
	fmt.Printf("The ID is: %v", id)
	err := r.DB.Where("id = ?", id).First(book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message" : " Couldn't get the book"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Book Found",
			"data": book,
		})
	return nil
}

func(r *Repository) DeleteBook(context *fiber.Ctx) error {
	book := models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "ID cannot be Empty"})
			return nil
	}

	err := r.DB.Delete(book, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Couldn't Delete Book"})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Book Deleted Successfully"})
	return nil
}