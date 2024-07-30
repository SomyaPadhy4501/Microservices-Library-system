package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SomyaPadhy4501/book-store/pkg/models"
	"github.com/SomyaPadhy4501/book-store/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

//Get all books controller

func GetBooks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	newBooks, err := models.GetAllBooks(userID)
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(newBooks)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//Get book by ID controller

func GetBookById(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal("error while Parsing:%w", err)
	}
	bookDetails, _ := models.GetBookById(ID, userID)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

//Creating a book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)

	CreateBook.UserID = userID
	b := CreateBook.CreateABook()

	res, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal("Error while parsing %w", err)
	}
	if err := models.DeleteBook(ID, userID); err != nil {
		http.Error(w, "Error deleting book: "+err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "Book deleted successfully"}
	res, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//Update by ID

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal("error while parsing %w", err)
	}
	bookDetails, err := models.GetBookById(ID, userID)
	if err != nil {
		http.Error(w, "Book not found or you don't have permission to update it", http.StatusNotFound)
		return
	}
	if updateBook.Name != bookDetails.Name && updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	if err := models.UpdateBook(bookDetails); err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(bookDetails)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
