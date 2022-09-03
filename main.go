package main

import (
	"database/sql"
	_ "database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"gangashinienovels_backend/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)


var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		logFatal(err)
	}
}

func main() {
	pgURL, err := pq.ParseURL(os.Getenv("PG_URL"))
	logFatal(err)
	db, err = sql.Open("postgres", pgURL)
	logFatal(err)
	err = db.Ping()
	logFatal(err)
	log.Println(pgURL)
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	books = []models.Book{}

	rows, err := db.Query("Select * from books")
	logFatal(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.Description, &book.Thumbnail)
		logFatal(err)
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	rows := db.QueryRow("select * from books where id=$1", params["id"])
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.Description, &book.Thumbnail)
	logFatal(err)
	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)
	err := db.QueryRow("insert into books (title , author, year ,description, thumbnail) values ($1,$2,$3,$4,$5) RETURNING id",
		book.Title, book.Author, book.Year, book.Description, book.Thumbnail).Scan(&bookID)
	logFatal(err)
	json.NewEncoder(w).Encode(bookID)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	result, err := db.Exec("update books set title=$1,author=$2,year=$3,description=$4,thumbnail=$5 where id=$6 RETURNING id", book.Title, book.Author, book.Year, book.Description, book.Thumbnail, book.ID)
	rowsUpdated, err := result.RowsAffected()
	logFatal(err)
	//return no of rows updated
	json.NewEncoder(w).Encode(rowsUpdated)
}
func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Exec("delete from books where id=$1", params["id"])
	logFatal(err)
	rowsDeleted, err := result.RowsAffected()
	logFatal(err)
	json.NewEncoder(w).Encode(rowsDeleted)
}
