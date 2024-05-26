package models

import (
	"github.com/jinzhu/gorm"
	"github.com/swarnendu19/Book-management-Store/pkg/config"
)

// Importing necessary packages for defining database models and establishing database connection

// Declaring a package-level variable to hold the database connection
var db *gorm.DB

// Defining the structure of the Book model with corresponding database fields
// Which including Name Author Publication

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name" `
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Initializing the database connection when the models package is imported
//
// Performing auto migration to create necessary tables if they don't exist
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook creates a new book record in the database.
// Records the new book entry and returns a pointer to the created book object.

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

// GetAllBooks retrieves all book records from the database.
// Fetches all book records and returns them as a slice of Book objects.

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books

}

// GetBookById retrieves a book record from the database based on the provided ID.
// Searches for a book record with the given ID and returns a pointer to the book object and a reference to the database transaction.

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

// DeleteBook deletes a book record from the database based on the provided ID.
// Deletes the book record with the given ID and returns the deleted book object.

func DeleteBook(id int64) Book {
	var deletedBook Book
	db.Where("ID=?", id).Delete(&deletedBook)
	return deletedBook
}
