package user

import (
	"book-api-mvc/lib/database"
	"book-api-mvc/models"
	"github.com/labstack/echo"
	"github.com/tegarap/simple-jsend"
	"net/http"
	"strconv"
)

func LoginUsersController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	users, err := database.LoginUser(&user)
	if err != nil {
		//return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("login failed"))
		return c.JSON(http.StatusBadRequest, common.ResponseFail("login failed"))
	}

	return c.JSON(http.StatusOK, common.ResponseSuccess(users))
}

func GetAllUserController(c echo.Context) error {
	users, err := database.GetAllUser()
	if err != nil {
		//return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("failed while fetching data"))
		return c.JSON(http.StatusBadRequest, common.ResponseFail("failed while fetching data"))
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess(users))
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		//return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("invalid param"))
		return c.JSON(http.StatusBadRequest, common.ResponseFail("invalid param"))
	}

	user, err := database.GetUser(id)
	if err != nil {
		//return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("id not found"))
		return c.JSON(http.StatusBadRequest, common.ResponseFail("id not found"))
	}

	return c.JSON(http.StatusOK, common.ResponseSuccess(user))
}

func CreateUserController(c echo.Context) error {
	usrInput := models.User{}
	c.Bind(&usrInput)

	if usrInput.Name == "" || usrInput.Password == "" || usrInput.Email == "" {
		return c.JSON(http.StatusBadRequest, common.ResponseFail("field are required"))
	}

	user, err := database.CreateUser(&usrInput)
	if err != nil {
		//return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("failed to create user"))
		return c.JSON(http.StatusBadRequest, common.ResponseFail("failed to create user"))
	}

	return c.JSON(http.StatusOK, common.ResponseSuccess(user))
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("invalid param"))
		return c.JSON(http.StatusBadRequest, common.ResponseFail("invalid parameter"))
	}
	_, err = database.GetUser(id)

	if err != nil {
		//return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("id not found"))
		return c.JSON(http.StatusBadRequest, common.ResponseFail("id not found"))
	} else {
		_, err = database.DeleteUser(id)
		if err != nil {
			//return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("cannot delete user"))
			return c.JSON(http.StatusBadRequest, common.ResponseFail("cannot delete user"))
		}
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess("success delete user with id: "+c.Param("id")))
}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("invalid param"))
		return c.JSON(http.StatusBadRequest, common.ResponseFail("invalid parameter"))
	}

	user, err := database.GetUser(id)
	if err != nil {
		//return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("id not found"))
		return c.JSON(http.StatusBadRequest, common.ResponseFail("id not found"))
	}

	var newUser models.User
	c.Bind(&newUser)

	upUser, err := database.UpdateUser(user, &newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ResponseFail("cannot delete user"))
	}

	return c.JSON(http.StatusOK, common.ResponseSuccess(upUser))
}
