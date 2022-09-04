package main

import (
	"database/sql"
	_ "database/sql"
	"gangashinienovels_backend/services/book/controllers"
	"gangashinienovels_backend/services/book/driver"
	"gangashinienovels_backend/services/book/models"

	"log"
	"net/http"

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
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}
