package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID          int    `json:id`
	Title       string `json:title`
	Author      string `json:author`
	Year        string `json:year`
	Description string `json:description`
	Thumbnail   string `json:thumbnail`
}

var books []Book

func main() {
	router := mux.NewRouter()
	books = append(books, Book{ID: 1, Title: "A", Author: "A", Year: "2022", Description: "sample book A description", Thumbnail: "www.google.lk"},
		Book{ID: 2, Title: "B", Author: "A", Year: "2022", Description: "sample book B description", Thumbnail: "www.google.lk"},
		Book{ID: 3, Title: "C", Author: "A", Year: "2022", Description: "sample book C description", Thumbnail: "www.google.lk"})
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	//log.Println("Get all books")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get a book")
}
func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add a books")
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update a books")
}
func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove a books")
}
