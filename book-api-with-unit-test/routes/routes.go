package routes

import (
	"book-api-mvc/controllers/book"
	"book-api-mvc/controllers/user"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	//jwtAuth := e.Group("")
	//jwtAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))
	//
	//e.POST("/login", user.LoginUsersController)
	//// ------------------------------------------------------------------
	//// USER
	//// ------------------------------------------------------------------
	//jwtAuth.GET("/users", user.GetAllUserController)
	//jwtAuth.GET("/users/:id", user.GetUserController)
	//e.POST("/users", user.CreateUserController)
	//jwtAuth.DELETE("/users/:id", user.DeleteUserController)
	//jwtAuth.PUT("/users/:id", user.UpdateUserController)
	//// ------------------------------------------------------------------
	//// BOOK
	//// ------------------------------------------------------------------
	//e.GET("/booksz", book.GetAllBookController)
	//e.GET("/booksz/:id", book.GetBookController)
	//e.POST("/booksz", book.CreateBookController)
	//jwtAuth.DELETE("/books/:id", book.DeleteBookController)
	//jwtAuth.PUT("/books/:id", book.UpdateBookController)



	// FOR TEST ONLY
	e.POST("/login", user.LoginUsersController)
	// ------------------------------------------------------------------
	// USER
	// ------------------------------------------------------------------
	e.GET("/users", user.GetAllUserController)
	e.GET("/users/:id", user.GetUserController)
	e.POST("/users", user.CreateUserController)
	e.DELETE("/users/:id", user.DeleteUserController)
	e.PUT("/users/:id", user.UpdateUserController)
	// ------------------------------------------------------------------
	// BOOK
	// ------------------------------------------------------------------
	e.GET("/booksz", book.GetAllBookController)
	e.GET("/booksz/:id", book.GetBookController)
	e.POST("/booksz", book.CreateBookController)
	e.DELETE("/books/:id", book.DeleteBookController)
	e.PUT("/books/:id", book.UpdateBookController)

	return e
}
