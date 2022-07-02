package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// User is the main entity
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserHandler analyses the request and delegates to the proper function
func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/user")
	id, err := strconv.Atoi(sid)
	if sid != "" && err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "I am not sure what you are looking for...")
		return
	}

	switch {
	case r.Method == "GET" && id > 0:
		getUserByID(w, r, id)
	case r.Method == "GET":
		getUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "I am not sure what you are looking for...")
	}
}

func getUserByID(w http.ResponseWriter, r *http.Request, id int) {
	db := openDBConnection(w)
	defer db.Close()

	var user User
	db.QueryRow("select id, name from users where id = ?", id).Scan(&user.ID, &user.Name)

	if user.ID <= 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "This guy does not exist!")
		return
	}

	jsonUser, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonUser))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	db := openDBConnection(w)
	defer db.Close()

	rows, _ := db.Query("select id, name from users")

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	jsonUsers, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonUsers))
}
