package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // third party router
)

const PORT int = 8000

// =========  Models/Structs ========= 
type Book struct {
	ID      string  `json:"id"`
	Isbn    string  `json:"isbn"`
	Title   string  `json:"title"`
	Author  *Author `json:"author"`
}

type Author struct {
	Firstname  string  `json:"fname"`
	Lastname   string  `json:"lname"`
}

// =========  Mock Data as slice Book struct ========= 
var books []Book

// 	========= Functions ========= 

// GetBooks
func getBooks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type", "Appliaction/json")	
  json.NewEncoder(w).Encode(books)
}

// Get one book by ID
func getBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type", "Appliaction/json")
  params := mux.Vars(r) // Get params
  
  for _, item := range books {

	if item.ID == params["id"] {
	   json.NewEncoder(w).Encode(item)
	   return
	}
  }

  json.NewEncoder(w).Encode(&Book{})
}
// createBook
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Appliaction/json")
    
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
  
}
// updateBook
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")

	params := mux.Vars(r) // Get params
	
	for idx, item := range books {
	  if item.ID == params["id"] {
		  books = append(books[:idx], books[idx+1:]...)
		 
		  var book Book

		  _ = json.NewDecoder(r.Body).Decode(&book)
		  book.ID = params["id"]
		  books = append(books, book)
		  json.NewEncoder(w).Encode(book)
		  
		  return 
	   }
	}
  
	json.NewEncoder(w).Encode(books)
  
}
// deleteBook
func deleteBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type", "Application/json")

  params := mux.Vars(r) // Get params
  
  for idx, item := range books {
	if item.ID == params["id"] {
        books = append(books[:idx], books[idx+1:]...)
		break
 	}
  }

  json.NewEncoder(w).Encode(books)
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