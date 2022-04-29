package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guhkun13/tutorial/freeCodeCamp/03-bookstore-api/pkg/models"
	"github.com/guhkun13/tutorial/freeCodeCamp/03-bookstore-api/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("err while parsing")
	}
	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)

	// fmt.Println(book)
	// fmt.Println("HALO")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("err while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var data = &models.Book{}
	utils.ParseBody(r, data)

	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("err while parsing")
	}
	book, db := models.GetBookById(ID)

	if book.Name != "" {
		book.Name = data.Name
	}

	if book.Author != "" {
		book.Author = data.Author
	}

	if book.Publication != "" {
		book.Publication = data.Publication
	}

	db.Save(&book)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}