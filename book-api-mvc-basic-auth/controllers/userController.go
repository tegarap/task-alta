package controllers

import (
	"book-api-mvc/controllers/util"
	"book-api-mvc/lib/database"
	"book-api-mvc/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, util.Response("success get all user", users))
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, e := database.GetUser(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, util.Response("success get user", user))
}

func CreateUserController(c echo.Context) error {
	usrInput := models.User{}
	err := c.Bind(&usrInput)
	if err != nil {
		return err
	}
	if usrInput.Name == "" || usrInput.Email == "" || usrInput.Password == "" {
		return c.JSON(http.StatusOK, util.Response("field are required", nil))
	}
	user, e := database.CreateUser(&usrInput)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, util.Response("success create user", user))
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	_, e := database.GetUser(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	} else {
		_, err := database.DeleteUser(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, e.Error())
		}
	}
	return c.JSON(http.StatusOK, util.Response("success delete user", nil))
}

func UpdateUserController(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	usr := models.User{}
	user, err := database.UpdateUser(id, &usr)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	return c.JSON(http.StatusOK, util.Response("success update user", user))
}
