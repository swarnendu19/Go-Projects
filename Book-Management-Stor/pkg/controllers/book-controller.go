package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/swarnendu19/Book-management-Store/pkg/models"
	"github.com/swarnendu19/Book-management-Store/pkg/utils"
)

// Importing necessary packages for handling HTTP requests, encoding/decoding JSON, and routing
// Importing utility functions and database models from the application packages

// Declaring a package-level variable to hold a new Book object

var newBook models.Book

// GetBook retrieves all books from the database and returns them as JSON response.
// Fetches all book records, marshals them into JSON format, and sends them as a response.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, err := json.Marshal(newBook)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "pkalication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBookById retrieves a specific book by ID from the database and returns it as JSON response.
// Extracts the book ID from the request parameters, retrieves the corresponding book record, and sends it as a response.

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bookId"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		panic(err)
	}
	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook creates a new book record in the database based on the JSON payload in the request body.
// Parses the request body into a Book object, creates a new book record, and sends the created book as a response.

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook deletes a book record from the database based on the provided ID.
// Extracts the book ID from the request parameters, deletes the corresponding book record, and sends the deleted book as a response.

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, _ := strconv.ParseInt(bookId, 0, 0)
	deletedBook := models.DeleteBook(id)
	res, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook updates an existing book record in the database based on the provided ID and JSON payload in the request body.
// Parses the request body into a Book object, retrieves the existing book record, updates its fields with the new data, saves the changes, and sends the updated book as a response.

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["bookId"], 0, 0)
	bookDetails, db := models.GetBookById(id)
	if book.Name != "" {
		bookDetails.Name = book.Name
	}
	if book.Author != "" {
		bookDetails.Author = book.Author
	}
	if book.Publication != "" {
		bookDetails.Publication = book.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
