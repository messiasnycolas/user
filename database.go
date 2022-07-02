package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var dbVars = loadEnv("DB_USERNAME", "DB_PASSWORD", "DB_NAME")
var dbUsername, dbPassword, dbName = dbVars[0], dbVars[1], dbVars[2]
var dbUrl string = fmt.Sprintf("%v:%v@/%v", dbUsername, dbPassword, dbName)

func openDBConnection(w http.ResponseWriter) (db *sql.DB) {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Oops! There is something wrong with my DB.")
	}
	return
}
