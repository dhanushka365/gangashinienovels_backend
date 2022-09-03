package main

import (
	"database/sql"
	_ "database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
	_ "github.com/subosito/gotenv"

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
var db *sql.DB

func init(){
	gotenv.Load()
}

func logFatal(err error){
	if err !=nil{
	logFatal(err)
	}
}

func main() {
	pgURL ,err :=pq.ParseURL(os.Getenv("PG_URL"))
	logFatal(err)
	db, err = sql.Open("postgres", pgURL)
	logFatal(err)
	err = db.Ping()
	logFatal(err)
	log.Println(pgURL)
	router := mux.NewRouter()
	books = append(books, Book{ID: 1, Title: "A", Author: "A", Year: "2022", Description: "sample book A description", Thumbnail: "www.google.lk"},
		Book{ID: 2, Title: "B", Author: "A", Year: "2022", Description: "sample book B description", Thumbnail: "www.google.lk"},
		Book{ID: 3, Title: "C", Author: "A", Year: "2022", Description: "sample book C description", Thumbnail: "www.google.lk"})
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	//log.Println("Get all books")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	//log.Println("Get a book")
	params := mux.Vars(r)
	//log.Println(params)
	i, _ := strconv.Atoi(params["id"])
	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}
func addBook(w http.ResponseWriter, r *http.Request) {
	//log.Println("Add a books")
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	//log.Println("Update a books")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID{
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(books)
}
func removeBook(w http.ResponseWriter, r *http.Request) {
	//log.Println("Remove a books")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range books{
		if item.ID == id{
			books =append(books[:1],books[i+1:]... )
		}
	}
	json.NewEncoder(w).Encode(books)
}
