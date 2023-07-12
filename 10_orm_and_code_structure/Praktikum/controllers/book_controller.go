package controllers

import (
	"net/http"

	"github.com/dedit-hs/go_dedit-hery-suprastyo/10_orm_and_code_structure/Praktikum/config"
	"github.com/dedit-hs/go_dedit-hery-suprastyo/10_orm_and_code_structure/Praktikum/models"
	"github.com/labstack/echo/v4"
)

func GetBooksController(e echo.Context) error {
	var books []models.Book
	err := config.DB.Find(&books)

	if err.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Error",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, models.Response{
		Message: "Success",
		Data:    books,
	})
}

func GetBookController(e echo.Context) error {
	id := e.Param("id")
	var book models.Book

	err := config.DB.First(&book, id)

	if err.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Error",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, models.Response{
		Message: "Success",
		Data:    book,
	})
}

func AddBookController(e echo.Context) error {
	var book models.Book
	e.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}
	return e.JSON(http.StatusCreated, models.Response{
		Message: "Success",
		Data:    book,
	})
}

func DeleteBookController(e echo.Context) error {
	id := e.Param("id")
	var book models.Book
	found := config.DB.First(&book, id)

	if found.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	deleteBook := config.DB.Delete(&book)
	if deleteBook.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, models.Response{
		Message: "Success",
		Data:    nil,
	})
}

func UpdateBookController(e echo.Context) error {
	id := e.Param("id")
	var book models.Book

	found := config.DB.First(&book, id)

	if found.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	e.Bind(&book)
	updatedBook := config.DB.Save(&book)

	if updatedBook.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, models.Response{
		Message: "Success",
		Data:    book,
	})
}
