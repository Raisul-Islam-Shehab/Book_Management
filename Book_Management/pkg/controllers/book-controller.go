package controllers

import (
	"Book_Management/pkg/models"
	"Book_Management/pkg/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(c echo.Context) error {
	newBooks := models.GetAllBooks()
	return c.JSON(http.StatusOK, newBooks)
}

func GetBookById(c echo.Context) error {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error when parsing")
	}

	bookDetails, _ := models.GetBookById(ID)

	return c.JSON(http.StatusOK, bookDetails)
}

func CreateBook(c echo.Context) error {
	// CreateBook := &models.Book{}

	CreateBook := new(models.Book)
	utils.ParseBody(c, CreateBook)
	b := CreateBook.CreateBook()
	return c.JSON(http.StatusOK, b)
}

func DeleteBookById(c echo.Context) error {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error when parsing")
	}
	book := models.DeleteBookById(ID)
	return c.JSON(http.StatusOK, book)
}

func UpdateBookById(c echo.Context) error {
	bookId := c.Param("bookId")
	updateBook := new(models.Book)
	utils.ParseBody(c, updateBook)

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error when parsing")
	}

	bookDetails, db := models.GetBookById(ID)
	if bookDetails.Name != updateBook.Name {
		bookDetails.Name = updateBook.Name
	}

	if bookDetails.Author != updateBook.Author {
		bookDetails.Author = updateBook.Author
	}

	if bookDetails.Publication != updateBook.Publication {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	return c.JSON(http.StatusOK, bookDetails)
}
