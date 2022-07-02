package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dbUsername = os.Getenv("DB_USERNAME")
var dbPassword = os.Getenv("DB_PASSWORD")
var dbName = os.Getenv("DB_NAME")
var dbUrl string = fmt.Sprintf("%v:%v@/%v", dbUsername, dbPassword, dbName)

func openDBConnection(w http.ResponseWriter) (db *sql.DB) {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Oops! There is something wrong with my DB.")
	}
	return
}
