package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// UserRoutes analyses the request and delegates to the proper function
func UserRoutes(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/user/")
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
	case r.Method == "POST":
		createUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "I am not sure what you are looking for...")
	}
}
