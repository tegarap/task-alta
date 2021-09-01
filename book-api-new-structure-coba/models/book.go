package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}


type GormBookModel struct {
	db *gorm.DB
}

func NewBookModel(db *gorm.DB) *GormBookModel {
	return &GormBookModel{db: db}
}

// Interface Book

type BookModel interface {
	GetAll() ([]Book, error)
	Get(bookId int) (Book, error)
	Insert(Book) (Book, error)
	Delete(bookId int) (Book, error)
	Edit(book Book, bookId int) (Book, error)
}

func (m *GormBookModel) GetAll() ([]Book, error) {
	var books []Book
	if err := m.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (m *GormBookModel) Get(bookId int) (Book, error) {
	var book Book
	if err := m.db.First(&book, bookId).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (m *GormBookModel) Insert(book Book) (Book, error) {
	if err := m.db.Create(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (m *GormBookModel) Delete(bookId int) (Book, error) {
	var book Book
	if err := m.db.First(&book, bookId).Error; err != nil {
		return book, err
	}
	if err := m.db.Delete(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (m *GormBookModel) Edit(bookTemp Book, bookId int) (Book, error) {
	var book Book
	if err := m.db.First(&book, bookId).Error; err != nil {
		return book, err
	}
	if bookTemp.Title != "" {
		book.Title = bookTemp.Title
	}
	if bookTemp.Author != "" {
		book.Author = bookTemp.Author
	}
	if bookTemp.Publisher != "" {
		book.Publisher = bookTemp.Publisher
	}
	if err := m.db.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}
