package controllers

import (
	"database/sql"
	"encoding/json"
	"gangashinienovels_backend/services/book/models"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		logFatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)

		rows := db.QueryRow("select * from books where id=$1", params["id"])
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.Description, &book.Thumbnail)
		logFatal(err)
		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int

		json.NewDecoder(r.Body).Decode(&book)
		err := db.QueryRow("insert into books (title , author, year ,description, thumbnail) values ($1,$2,$3,$4,$5) RETURNING id",
			book.Title, book.Author, book.Year, book.Description, book.Thumbnail).Scan(&bookID)
		logFatal(err)
		json.NewEncoder(w).Encode(bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)
		result, err := db.Exec("update books set title=$1,author=$2,year=$3,description=$4,thumbnail=$5 where id=$6 RETURNING id", book.Title, book.Author, book.Year, book.Description, book.Thumbnail, book.ID)
		rowsUpdated, err := result.RowsAffected()
		logFatal(err)
		//return no of rows updated
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
	
		result, err := db.Exec("delete from books where id=$1", params["id"])
		logFatal(err)
		rowsDeleted, err := result.RowsAffected()
		logFatal(err)
		json.NewEncoder(w).Encode(rowsDeleted)
	}
}


