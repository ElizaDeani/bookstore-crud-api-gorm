package bookstoreroutes

import (
	"post-test-mikti/pkg/controllers/bookcontrollers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterBookRoutes(router *mux.Router, db *gorm.DB) {
	// Rute dasar untuk buku
	booksRouter := router.PathPrefix("/api/books").Subrouter()

	booksRouter.HandleFunc("/", bookcontrollers.GetAllBooks(db)).Methods("GET")
	booksRouter.HandleFunc("/{id}", bookcontrollers.GetBookByID(db)).Methods("GET")
	booksRouter.HandleFunc("/", bookcontrollers.CreateBook(db)).Methods("POST")
	booksRouter.HandleFunc("/{id}", bookcontrollers.UpdateBook(db)).Methods("PUT")
	booksRouter.HandleFunc("/{id}", bookcontrollers.DeleteBook(db)).Methods("DELETE")
}
