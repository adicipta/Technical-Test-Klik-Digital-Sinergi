package controllers

import (
	"net/http"
	"strconv"

	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/lib/database"
	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/models"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	db database.BookModel
}

func NewBookController(db database.BookModel) *BookController {
	return &BookController{db: db}
}

func (bc *BookController) GetBooks(c echo.Context) error {
	books, err := bc.db.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}
	var result models.Book_Response
	var resultSlice []models.Book_Response
	for _, v := range books {
		result.ID = v.ID
		result.Name = v.Name
		result.Description = v.Description
		result.Quantity = v.Quantity
		resultSlice = append(resultSlice, result)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   resultSlice,
	})
}

func (bc *BookController) AddBook(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Input")
	}
	book, err := bc.db.AddBook(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"data": map[string]interface{}{
			"id":          book.ID,
			"name":        book.Name,
			"description": book.Description,
		},
	})
}

func (bc *BookController) EditBooks(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}
	var newBook models.Book
	c.Bind(&newBook)
	newBook, err = bc.db.EditBookById(id, newBook)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data": map[string]interface{}{
			"id":          newBook.ID,
			"name":        newBook.Name,
			"description": newBook.Description,
		},
	})
}

func (bc *BookController) DeleteBooks(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Id")
	}

	err = bc.db.DeleteBookById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}
