package models

import (
	"github.com/SomyaPadhy4501/book-store/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//Making book data

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	UserID      string `json:"user_id"`
}

// Initialization

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//Creating the data base function
// Create Book function

func (b *Book) CreateABook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// All books

func GetAllBooks(userID string) ([]Book, error) {
	db := config.GetDB()
	var books []Book
	if err := db.Where("user_id = ?", userID).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

//Get Book by ID

func GetBookById(Id int64, userID string) (*Book, error) {
	var getBook Book
	if err := db.Where("id = ? AND user_id = ?", Id, userID).First(&getBook).Error; err != nil {
		return nil, err
	}
	return &getBook, nil
}

//Delete book

func DeleteBook(ID int64, userID string) error {
	db := config.GetDB()
	if err := db.Where("id = ? AND user_id = ?", ID, userID).Delete(&Book{}).Error; err != nil {
		return err
	}
	return nil
}

// Update book
func UpdateBook(book *Book) error {
	db := config.GetDB()
	if err := db.Save(book).Error; err != nil {
		return err
	}
	return nil
}
