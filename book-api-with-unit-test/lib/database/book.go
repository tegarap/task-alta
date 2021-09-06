package database

import (
	"book-api-mvc/config"
	"book-api-mvc/models"
)

func GetAllBook() (interface{}, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBook(id int) (interface{}, error) {
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func CreateBook(book *models.Book) (interface{}, error) {
	if err := config.DB.Create(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(id int) (interface{}, error) {
	var book models.Book
	if err := config.DB.Delete(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(book interface{}, newBook models.Book) (interface{}, error) {
	if err := config.DB.Model(book).Updates(newBook).Error; err != nil {
		return nil, err
	}
	return book, nil
}
