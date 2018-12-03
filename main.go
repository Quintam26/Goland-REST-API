package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

//Book Struct(model...like a es6 class)
type Book struct {
	ID 			string  `json:"id"`
	Isbn 		string  `json:"isbn"`
	Title 	string  `json:"title"`
	Author 	*Author `json:"author"`
}

//Author Struct
type Author struct {
	Firstname 	string	`json:"firstname"`
	Lasttname 	string	`json:"lastname"`
}

//Init books variable as a slice Book struct
var books []Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	//Loop through books and find id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) //Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

//Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	//Init the router
	router := mux.NewRouter()

	//Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "5423", Title: "Book One", Author: &Author{Firstname: "John", Lasttname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "1111", Title: "Book Two", Author: &Author{Firstname: "Steve", Lasttname: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "8654", Title: "Book Three", Author: &Author{Firstname: "Dave", Lasttname: "Macho"}})

	//Route  handlers/endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//Run server
	log.Fatal(http.ListenAndServe(":8000", router))
}