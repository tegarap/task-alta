package controllers

import (
	"book-api-mvc/controllers/common"
	"book-api-mvc/lib/database"
	"book-api-mvc/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewBadRequestResponse())
	}
	return c.JSON(http.StatusOK, common.Response("success login", users))
}

func GetUserDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewBadRequestResponse())
	}
	users, err := database.GetDetailUsers(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewBadRequestResponse())
	}
	return c.JSON(http.StatusOK, common.Response("success", users))
}

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewBadRequestResponse())
	}
	return c.JSON(http.StatusOK, common.Response("success get all user", users))
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewBadRequestResponse())
	}
	user, e := database.GetUser(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewBadRequestResponse())
	}
	return c.JSON(http.StatusOK, common.Response("success get user", user))
}

func CreateUserController(c echo.Context) error {
	usrInput := models.User{}
	err := c.Bind(&usrInput)
	if err != nil {
		return err
	}
	if usrInput.Name == "" || usrInput.Email == "" || usrInput.Password == "" {
		return c.JSON(http.StatusOK, common.Response("field are required", nil))
	}
	user, e := database.CreateUser(&usrInput)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewBadRequestResponse())
	}
	return c.JSON(http.StatusOK, common.Response("success create user", user))
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewBadRequestResponse())
	}
	_, e := database.GetUser(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewBadRequestResponse())
	} else {
		_, err := database.DeleteUser(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, common.NewBadRequestResponse())
		}
	}
	return c.JSON(http.StatusOK, common.Response("success delete user", nil))
}

func UpdateUserController(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewBadRequestResponse())
	}
	usr := models.User{}
	_, er := database.GetUser(id)
	er = c.Bind(&usr)
	if er != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}
	user, err := database.UpdateUser(id, &usr)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, common.NewBadRequestResponse())
	}
	return c.JSON(http.StatusOK, common.Response("success update user", user))
}
