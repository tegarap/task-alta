package controllers

import (
	"book-api-mvc/controllers/common"
	"book-api-mvc/lib/database"
	"book-api-mvc/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetAllBookController(c echo.Context) error {
	books, err := database.GetAllBook()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("failed while fetching data"))
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess(books))
}

func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("invalid param"))
	}
	book, err := database.GetBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("id not found"))
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess(book))
}

func CreateBookController(c echo.Context) error {
	usrInput := models.Book{}
	err := c.Bind(&usrInput)
	if err != nil {
		return err
	}
	if usrInput.Title == "" || usrInput.Author == "" || usrInput.Publisher == "" {
		return c.JSON(http.StatusBadRequest, common.ResponseFail("field are required"))
	}
	book, err := database.CreateBook(&usrInput)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("failed add book"))
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess(book))
}

func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("invalid param"))
	}
	_, err = database.GetBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("id not found"))
	} else {
		_, err := database.DeleteBook(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("cannot delete book"))
		}
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess("success delete book with id: "+c.Param("id")))
}

func UpdateBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("invalid param"))
	}

	book, err := database.GetBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("id not found"))
	}

	var newBook models.Book
	c.Bind(&newBook)

	upBook, err := database.UpdateBook(book, newBook)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("cannot delete book"))
	}

	return c.JSON(http.StatusOK, common.ResponseSuccess(upBook))
}
