package config

import (
	"book-api-mvc/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// initDB
	config := map[string]string{
		"DbUsername": "tegarap",
		"DbPassword": "t00r!Roo",
		"DbPort":     "3306",
		"DbHost":     "localhost",
		"DbName":     "crud_go",
	}
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["DbUsername"],
		config["DbPassword"],
		config["DbHost"],
		config["DbPort"],
		config["DbName"])
	var err error
	DB, err = gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate()  {
	err := DB.AutoMigrate(&models.User{}, &models.Book{})
	if err != nil {
		return
	}
}
