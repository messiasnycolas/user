package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// UserRoutes analyses the request and delegates to the proper function
func UserRoutes(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/user")
	sid = strings.ReplaceAll(sid, "/", "")
	id, err := strconv.ParseInt(sid, 10, 64)

	if sid != "" && err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "I am not sure what you are looking for...")
		return
	}

	switch {
	case r.Method == "DELETE" && id > 0:
		deleteUser(w, r, id)
	case r.Method == "PUT" && id > 0:
		updateUser(w, r, id)
	case r.Method == "GET" && id > 0:
		getUserByID(w, r, id)
	case r.Method == "GET":
		getUsers(w, r)
	case r.Method == "POST":
		createUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "I am not sure what you are looking for...")
	}
}
