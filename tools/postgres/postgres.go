package postgres

import (
	"database/sql"
	"fmt"
	"log"

	//pg is needed for Postgres drivers
	_ "github.com/lib/pq"
)

//const - database connection constants
const (
	Host     string = "localhost"
	Port     string = "5432"
	User     string = "postgres"
	Password string = "admin"
	Database string = "postgres"
)

//Dbconn - postgres connection string
var Dbconn = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", User, Password, Database, Host, Port)

const appUser string = "bgrudt"

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", Dbconn)
	if err != nil {
		log.Fatal(err)
	}
}
