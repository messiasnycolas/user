package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var dburl string = "root:123456@/golearning"

func openDBConnection(w http.ResponseWriter) (db *sql.DB) {
	db, err := sql.Open("mysql", dburl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Oops! There is something wrong with my DB.")
	}
	return
}
