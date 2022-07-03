package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// User is the main entity
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type IDResponse = struct {
	ID int64 `json:"id"`
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

func createUser(w http.ResponseWriter, r *http.Request) {
	db := openDBConnection(w)
	defer db.Close()

	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	if newUser.Name == "" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "User must have a name!")
		return
	}

	res, err := db.Exec("insert into users(name) values(?)", newUser.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error while creating user!")
		return
	}

	lastInsertId, _ := res.LastInsertId()
	response := IDResponse{lastInsertId}
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResponse))
}

func updateUser(w http.ResponseWriter, r *http.Request, id int) {
	db := openDBConnection(w)
	defer db.Close()

	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)

	_, err := db.Exec("update users set name=? where id=?", newUser.Name, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error while updating user!")
		return
	}

	fmt.Fprintf(w, "Successful operation.")
}

func deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	db := openDBConnection(w)
	defer db.Close()

	_, err := db.Exec("delete from users where id=?", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error while deleting user!")
		return
	}

	fmt.Fprintf(w, "Successful operation.")
}
