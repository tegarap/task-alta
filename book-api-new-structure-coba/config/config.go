package config

import (
	"book-api-mvc/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func DatabaseConnection() *gorm.DB {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	//connString := fmt.Sprintf("tegarap:t00r!Roo@tcp(localhost:3306)/crud_go?charset=utf8mb4&parseTime=True&loc=Local")

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DbUserMigration(db)
	DbBookMigration(db)

	return db
}

func DbUserMigration(db *gorm.DB)  {
	db.AutoMigrate(models.User{})
}

func DbBookMigration(db *gorm.DB)  {
	db.AutoMigrate(models.Book{})
}