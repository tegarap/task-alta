package database

import (
	"book-api-mvc/config"
	"book-api-mvc/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book
	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBook(id int) (interface{}, error) {
	book := models.Book{}
	if e := config.DB.First(&book, id).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func CreateBook(book *models.Book) (interface{}, error) {
	if e := config.DB.Create(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func DeleteBook(id int) (interface{}, error) {
	book := models.Book{}
	if e := config.DB.Delete(&book, id).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func UpdateBook(id int, input *models.Book) (interface{}, error) {
	book := models.Book{}
	e := config.DB.First(&book, id)
	if e.Error != nil {
		return nil, e.Error
	}
	if e.RowsAffected > 0 {
		if input.Title != "" {
			book.Title = input.Title
		}
		if input.Author != "" {
			book.Author = input.Author
		}
		if input.Publisher != "" {
			book.Publisher = input.Publisher
		}
		config.DB.Save(&book)
		return &book, nil
	}
	return nil, e.Error
}