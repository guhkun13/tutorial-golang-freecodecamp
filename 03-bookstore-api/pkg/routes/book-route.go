package routes

import (
	"github.com/gorilla/mux"
	"github.com/guhkun13/tutorial/freeCodeCamp/03-bookstore-api/pkg/controllers"
)

var RegisterBookStoreRoutes = func( router *mux.Router) {
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}