package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bobsar0/go-bookstore/pkg/models"
	"github.com/bobsar0/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()

	res, _ := json.Marshal(newBooks) //convert to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ")
	}

	bookDetails, _ := models.GetBookById(Id)

	res, _ := json.Marshal(bookDetails) //convert to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}

	utils.ParseBody(r, book)

	newBook := book.CreateBook()

	res, _ := json.Marshal(newBook) //convert to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ")
	}

	book := models.DeleteBook(Id)

	res, _ := json.Marshal(book) //convert to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}

	utils.ParseBody(r, book)

	// get id to update from request
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ")
	}

	// get book to update
	bookDetails, db := models.GetBookById(Id)

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

	res, _ := json.Marshal(bookDetails) //convert to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
