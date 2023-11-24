package main

import (
	"database/sql"
	"fmt"
	// "os"

	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB // created outside to make it global.

var (
	HOST = "localhost"
	PORT = 5432
	USER = "postgres"
	DBNAME = "postgres"
	PASSWORD = "postgres"
)

// make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase() {
	// set up postgres sql to open it.
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		HOST, PORT, USER, DBNAME, PASSWORD)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ", errSql)
		panic(errSql)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
	}
}
