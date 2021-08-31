package routes

import (
	"book-api-mvc/constants"
	c "book-api-mvc/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	e.POST("/login", c.LoginUsersController)
	// ------------------------------------------------------------------
	// USER
	// ------------------------------------------------------------------
	r.GET("/users", c.GetUsersController)
	r.GET("/users/:id", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	r.DELETE("/users/:id", c.DeleteUserController)
	r.PUT("/users/:id", c.UpdateUserController)
	r.GET("/users/:id", c.GetUserDetailController)
	// ------------------------------------------------------------------
	// BOOK
	// ------------------------------------------------------------------
	e.GET("/books", c.GetBooksController)
	e.GET("/books/:id", c.GetBookController)
	r.POST("/books", c.CreateBookController)
	r.DELETE("/books/:id", c.DeleteBookController)
	r.PUT("/books/:id", c.UpdateBookController)

	return e
}