package routes

import (
	"github.com/SomyaPadhy4501/book-store/pkg/controllers"
	"github.com/gorilla/mux"
)

func StoreRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
