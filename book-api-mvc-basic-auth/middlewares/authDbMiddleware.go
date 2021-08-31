package middlewares

import (
	"book-api-mvc/config"
	"book-api-mvc/models"
	"github.com/labstack/echo"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	db := config.DB
	user := models.User{}
	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
