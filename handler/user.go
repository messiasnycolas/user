package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/messiasnycolas/user/model"
)

type IDResponse = struct {
	ID int64 `json:"id"`
}

func getUserByID(w http.ResponseWriter, r *http.Request, id int64) {
	user, err := model.Get(id)
	if err != nil {
		// log error
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Unexpected error.")
		}
		return
	}

	jsonUser, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(jsonUser))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetAll()
	if err != nil {
		// log error
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unexpected error.")
		return
	}

	jsonUsers, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(jsonUsers))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	if user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User must have a name!")
		return
	}

	id, err := model.Insert(user)
	if err != nil || id == 0 {
		// log error
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unexpected error.")
		return
	}

	response := IDResponse{id}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(jsonResponse))
}

func updateUser(w http.ResponseWriter, r *http.Request, id int64) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	if user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User must have a name!")
		return
	}

	rowsAffected, err := model.Update(id, user)
	if err != nil || rowsAffected != 1 {
		// log error
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unexpected error.")
		return
	}

	fmt.Fprintf(w, "Successful operation.")
}

func deleteUser(w http.ResponseWriter, r *http.Request, id int64) {
	rowsAffected, err := model.Delete(id)
	if err != nil || rowsAffected != 1 {
		// log error
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unexpected error.")
		return
	}

	fmt.Fprintf(w, "Successful operation.")
}
