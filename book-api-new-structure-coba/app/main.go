package main

import (
	"book-api-mvc/api"
	bookController "book-api-mvc/api/controllers/book"
	userController "book-api-mvc/api/controllers/user"
	m "book-api-mvc/api/middlewares"
	"book-api-mvc/config"
	"book-api-mvc/models"
	"github.com/labstack/echo"
	"os"
)

func main() {
	// initialize database connection
	db := config.DatabaseConnection()

	// initialize user model
	userModel := models.NewUserModel(db)
	bookModel := models.NewBookModel(db)

	// initialize user controller
	newUserController := userController.NewController(userModel)
	newBookController := bookController.NewController(bookModel)

	// create echo http
	//e := routes.New()
	e := echo.New()

	//register API path and controller
	api.UserPath(e, newUserController)
	api.BookPath(e, newBookController)

	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":"+os.Getenv("SERV_PORT")))
}
