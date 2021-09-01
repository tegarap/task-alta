package api

import (
	"book-api-mvc/api/controllers/book"
	"book-api-mvc/api/controllers/user"
	"github.com/labstack/echo"
)

func UserPath(e *echo.Echo, userController *user.Controller) {
	// ------------------------------------------------------------------
	// Login & register
	// ------------------------------------------------------------------
	e.POST("/users/login", userController.LoginUserController)

	// ------------------------------------------------------------------
	// CRUD User
	// ------------------------------------------------------------------
	e.GET("/users", userController.GetAllUserController)
	e.GET("/users/:id", userController.GetUserController)
	e.POST("/users/users", userController.PostUserController)
	e.DELETE("/users/:id", userController.DeleteUserController)
	e.PUT("/users/:id", userController.EditUserController)
}

func BookPath(e *echo.Echo, bookController *book.Controller)  {
	// ------------------------------------------------------------------
	// CRUD Book
	// ------------------------------------------------------------------
	e.GET("/books", bookController.GetAllBookController)
	e.GET("/books/:id", bookController.GetBookController)
	e.POST("/books", bookController.PostBookController)
	e.DELETE("/books/:id", bookController.DeleteBookController)
	e.PUT("/books/:id", bookController.EditBookController)
}


//func New() *echo.Echo {
//	e := echo.New()
//	r := e.Group("")
//	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
//
//	e.POST("/login", controllers.LoginUsersController)
//	// ------------------------------------------------------------------
//	// USER
//	// ------------------------------------------------------------------
//	e.GET("/users", controllers.GetUsersController)
//	e.GET("/users/:id", controllers.GetUserController)
//	r.GET("/users/:id", controllers.GetUserDetailController)
//	e.POST("/users", controllers.CreateUserController)
//	e.DELETE("/users/:id", controllers.DeleteUserController)
//	e.PUT("/users/:id", controllers.UpdateUserController)
//	// ------------------------------------------------------------------
//	// BOOK
//	// ------------------------------------------------------------------
//	e.GET("/books", controllers.GetBooksController)
//	e.GET("/books/:id", controllers.GetBookController)
//	e.POST("/books", controllers.CreateBookController)
//	e.DELETE("/books/:id", controllers.DeleteBookController)
//	e.PUT("/books/:id", controllers.UpdateBookController)
//
//	return e
//}