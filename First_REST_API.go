package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Structure (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice book  struct
var books []Book

//Get All Books
func getbooks(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getbook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router) //Get params
	//Loop Through books and find id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new Book
func createbooks(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(router.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update
func updatebooks(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	params := mux.Vars(router)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			_ = json.NewDecoder(router.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete
func deletebooks(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//Init Router
	router := mux.NewRouter()

	// Example Data
	books = append(books, Book{ID: "1", Isbn: "456798", Title: "Book 1", Author: &Author{Firstname: "Will", Lastname: "Smith"}})

	books = append(books, Book{ID: "2", Isbn: "789456", Title: "Book 2", Author: &Author{Firstname: "Paul", Lastname: "Walker"}})

	//Route Handlers / Endpoints
	router.HandleFunc("/api/books", getbooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getbook).Methods("GET")
	router.HandleFunc("/api/books", createbooks).Methods("POST")
	router.HandleFunc("/api/books/{id}", updatebooks).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deletebooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router)) //Starting Server
}
