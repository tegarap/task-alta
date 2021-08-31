package routes

import (
	c "book-api-mvc/controllers"
	m "book-api-mvc/middlewares"
	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// User
	e.GET("/users", c.GetUsersController)
	e.GET("/users/:id", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	e.DELETE("/users/:id", c.DeleteUserController)
	e.PUT("/users/:id", c.UpdateUserController)

	//Book
	e.GET("/books", c.GetBooksController)
	e.GET("/books/:id", c.GetBookController)
	e.POST("/books", c.CreateBookController)
	e.DELETE("/books/:id", c.DeleteBookController)
	e.PUT("/books/:id", c.UpdateBookController)

	//middlewares
	eAuth := e.Group("")
	eAuth.Use(echoMid.BasicAuth(m.BasicAuthDB))

	eAuth.DELETE("/users/:id", c.DeleteUserController)
	eAuth.PUT("/users/:id", c.UpdateUserController)

	eAuth.DELETE("/books/:id", c.DeleteBookController)
	eAuth.PUT("/books/:id", c.UpdateBookController)

	return e
}