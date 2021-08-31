package middlewares

import (
	"github.com/labstack/echo"
	"os"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	if username == os.Getenv("USERNAME_LOGIN") && password == os.Getenv("PASSWORD_LOGIN") {
		return true, nil
	}
	return false, nil
}
