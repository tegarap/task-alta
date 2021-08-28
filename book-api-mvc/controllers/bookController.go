package controllers

import (
	"book-api-mvc/controllers/util"
	"book-api-mvc/lib/database"
	"book-api-mvc/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, util.Response("success get all book", books))
}

func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	book, e := database.GetBook(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, util.Response("success get book", book))
}

func CreateBookController(c echo.Context) error {
	usrInput := models.Book{}
	err := c.Bind(&usrInput)
	if err != nil {
		return err
	}
	if usrInput.Title == "" || usrInput.Author == "" || usrInput.Publisher == "" {
		return c.JSON(http.StatusOK, util.Response("field are required", nil))
	}
	book, e := database.CreateBook(&usrInput)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, util.Response("success create book", book))
}

func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	_, e := database.DeleteBook(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, util.Response("success delete book", nil))
}

func UpdateBookController(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	book := models.Book{}
	_, er := database.GetBook(id)
	er = c.Bind(&book)
	if er != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, er.Error())
	}
	upBook, err := database.UpdateBook(id, &book)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	return c.JSON(http.StatusOK, util.Response("success update book", upBook))
}
