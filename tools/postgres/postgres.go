package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", Dbconn)
	if err != nil {
		log.Fatal(err)
	}
}
