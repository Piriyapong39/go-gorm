package main

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Description string
	Price       uint
}

func CreateBook(db *gorm.DB, book *Book) error {
	result := db.Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetBook(db *gorm.DB, id int) (*Book, error) {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		return &Book{}, result.Error
	}
	return &book, nil
}

func UpdateBook(db *gorm.DB, book *Book) error {
	reslut := db.Save(&book)
	if reslut.Error != nil {
		return reslut.Error
	}
	return nil
}

func DeleteBook(db *gorm.DB, id int) error {
	var book Book
	result := db.Delete(&book, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
