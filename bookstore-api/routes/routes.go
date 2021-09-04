package routes

import (
	"book-api-mvc/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func New() *echo.Echo {
	e := echo.New()
	jwtAuth := e.Group("")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))

	e.POST("/login", controllers.LoginUsersController)
	// ------------------------------------------------------------------
	// USER
	// ------------------------------------------------------------------
	jwtAuth.GET("/users", controllers.GetAllUserController)
	jwtAuth.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	jwtAuth.DELETE("/users/:id", controllers.DeleteUserController)
	jwtAuth.PUT("/users/:id", controllers.UpdateUserController)
	// ------------------------------------------------------------------
	// BOOK
	// ------------------------------------------------------------------
	e.GET("/books", controllers.GetAllBookController)
	e.GET("/books/:id", controllers.GetBookController)
	jwtAuth.POST("/books", controllers.CreateBookController)
	jwtAuth.DELETE("/books/:id", controllers.DeleteBookController)
	jwtAuth.PUT("/books/:id", controllers.UpdateBookController)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))
	eJwt.GET("/users/:id", controllers.GetUserDetailControllers)

	return e
}
