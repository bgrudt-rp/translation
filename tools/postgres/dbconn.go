package postgres

import "fmt"

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
