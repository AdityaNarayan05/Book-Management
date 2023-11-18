package routes

import (
	"github.com/AdityaNarayan05/Book-Management/pkg/controllers"
	"github.com/gorilla/mux"
)

// RegisterBookStoreRoutes registers the routes related to the BookStore application.
var RegisterBookStoreRoutes = func(router *mux.Router) {
	// Handle POST requests to create a new book
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")

	// Handle GET requests to retrieve all books
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")

	// Handle GET requests to retrieve a book by its ID
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")

	// Handle PUT requests to update a book by its ID
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")

	// Handle DELETE requests to delete a book by its ID
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
