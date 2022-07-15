package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // third party router
)

const PORT int = 8000

// =========  Models/Structs ========= 
type Book struct {
	ID      string  `json: "id"`
	Isbn    string  `json: "isbn"`
	Title   string  `json: "title"`
	Author  *Author `json: "author"`
}

type Author struct {
	Firstname  string  `json:fname`
	Lastname   string  `json:lname`
}

// =========  Mock Data as slice Book struct ========= 
var books []Book

// 	========= Functions ========= 

// GetBooks
func getBooks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type", "Application/json")	
  json.NewEncoder(w).Encode(books)
}
// GetBook by ID
func getBook(w http.ResponseWriter, r *http.Request) {

  
}
// createBook
func createBook(w http.ResponseWriter, r *http.Request) {
  
}
// updateBook
func updateBook(w http.ResponseWriter, r *http.Request) {
  
}
// deleteBook
func deleteBook(w http.ResponseWriter, r *http.Request) {
  
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "45346", Title: "Book one", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "98792", Title: "Book two", Author: &Author{Firstname: "Mark", Lastname: "Tardelle"}})

	// Routes Handlers / Endpoints
	router.HandleFunc("/api/v1/books", getBooks).Methods("GET")
	router.HandleFunc("/api/v1/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/v1/books", createBook).Methods("POST")
	router.HandleFunc("/api/v1/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/v1/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server running...")

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(PORT), router))
}