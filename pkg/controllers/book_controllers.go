package controllers

import (
	"design_pattern/pkg/config"
	"design_pattern/pkg/models"
	"design_pattern/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBooks(c echo.Context) error {
	var books []models.Book
	result := config.DB.Find(&books)
	if result.Error != nil {
		utils.LogError(result.Error)
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	return c.JSON(http.StatusOK, books)
}

func GetBookByID(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		utils.LogError(result.Error)
		return c.JSON(http.StatusNotFound, result.Error)
	}

	return c.JSON(http.StatusOK, book)
}

func CreateBook(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		utils.LogError(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validasi title, author, dan price
	if err := utils.ValidateBookTitle(book.Title); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateBookAuthor(book.Author); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := utils.ValidatePrice(book.Price); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := config.DB.Create(&book)
	if result.Error != nil {
		utils.LogError(result.Error)
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	utils.LogInfo("Book created successfully")
	return c.JSON(http.StatusCreated, book)
}

func UpdateBook(c echo.Context) error {
	c.Param("id")
	var book models.Book
	if err := c.Bind(&book); err != nil {
		utils.LogError(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validasi title, author, dan price
	if err := utils.ValidateBookTitle(book.Title); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateBookAuthor(book.Author); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := utils.ValidatePrice(book.Price); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := config.DB.Save(&book)
	if result.Error != nil {
		utils.LogError(result.Error)
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	utils.LogInfo("Book updated successfully")
	return c.JSON(http.StatusOK, book)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	result := config.DB.Delete(&book, id)
	if result.Error != nil {
		utils.LogError(result.Error)
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	utils.LogInfo("Book deleted successfully")
	return c.NoContent(http.StatusNoContent)
}
