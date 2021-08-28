package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func main() {
	e := echo.New()
	// routing
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server
	e.Logger.Fatal(e.Start(":8000"))
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range users {
		if user.Id == id {
			u := new(User)
			if err := c.Bind(u); err != nil {
				return err
			}
			if u.Name != "" {
				users[i].Name = u.Name
			}
			if u.Email != "" {
				users[i].Email = u.Email
			}
			if u.Password != "" {
				users[i].Password = u.Password
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "update user with id " + c.Param("id") + " success",
				"user":     users[i],
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user with id: " + c.Param("id") + " not found",
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "delete user with id " + c.Param("id") + " success",
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user with id: " + c.Param("id") + " not found",
	})
}

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get user with id " + c.Param("id"),
				"user":     user,
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user with id: " + c.Param("id") + " not found",
	})
}
