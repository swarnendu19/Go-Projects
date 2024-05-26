package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/swarnendu19/Book-management-Store/pkg/controllers"
)

var RegisterBookStoreRoute = func(router *mux.Router) {
	// Make the routes of the Book Mnaement Store
	fmt.Println("")
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
