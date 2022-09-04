package driver

import (
	"database/sql"
	"os"
	"github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		logFatal(err)
	}
}

func ConnectDB() *sql.DB{
	pgURL, err := pq.ParseURL(os.Getenv("PG_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db

}

