package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welecome </h1>")
}

// Book Struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliation/json")
	json.NewEncoder(w).Encode(books)

}

// Get Single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliation/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

// Creat a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliation/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "appliation/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
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
func deleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "appliation/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
func main() {
	//http.HandleFunc("/", handlerFunc)
	//fmt.Println("Starting the server on :3000....")
	//http.ListenAndServe(":3000", nil)

	// Init Router
	r := mux.NewRouter()

	// Mock data - @todo - implatement db
	books = append(books, Book{ID: "1", Isbn: "44412", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "44413", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
