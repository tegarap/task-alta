package main

import (
	"fmt"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var (
	DB    *gorm.DB
	user  User
	users []User
)

type Config struct {
	DbUsername string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func init() {
	// initDB
	config := Config{
		DbUsername: "tegarap",
		DbPassword: "t00r!Roo",
		DbPort:     "3306",
		DbHost:     "localhost",
		DbName:     "crud_go",
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Automigration
	errMigration := DB.AutoMigrate(&User{})
	if errMigration != nil {
		return
	}
}

func main() {
	e := echo.New()
	//routing
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	//start and logging
	e.Logger.Fatal(e.Start(":8000"))
}

func UpdateUserController(c echo.Context) error {
	usr := new(User)
	id := c.Param("id")
	tx := DB.First(&user, id)
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, Response("data with id: "+id+" not found", nil))
	}
	if tx.RowsAffected > 0 {
		c.Bind(&usr)
		user.Name, user.Email, user.Password = usr.Name, usr.Email, usr.Password
		DB.Save(&user)
		return c.JSON(http.StatusOK, Response("user with id: "+id+" updated", nil))
	}
	return c.JSON(http.StatusOK, Response("user with id: "+id+" not found", nil))
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	tx := DB.Delete(&user, id)
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, Response("failed while fetching data", nil))
	}
	if tx.RowsAffected > 0 {
		return c.JSON(http.StatusOK, Response("success delete user", nil))
	}
	return c.JSON(http.StatusOK, Response("user with id: "+id+" not found", nil))
}

func CreateUserController(c echo.Context) error {
	usr := User{}
	c.Bind(&usr)
	if err := DB.Create(&usr).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, Response("cannot create user", nil))
	}
	return c.JSON(http.StatusOK, Response("success get user", usr))
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, Response("failed while fetching data", nil))
	}
	if errDB := DB.First(&user, id).Error; errDB != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, Response("user with id: "+c.Param("id")+" not found", nil))
	}
	return c.JSON(http.StatusOK, Response("success get user", user))
}

func GetUsersController(c echo.Context) error {
	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, Response("failed while fetching data", nil))
	}
	return c.JSON(http.StatusOK, Response("success get all user", users))
}

func Response(m string, r interface{}) map[string]interface{} {
	if r == nil {
		return map[string]interface{}{"message": m}
	}
	return map[string]interface{}{"message": m, "result": r}
}