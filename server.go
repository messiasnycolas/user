package main

import (
	"fmt"
	"log"
	"net/http"
)

var port int = 3000

func main() {
	http.HandleFunc("/user", UserHandler)
	log.Printf("Server listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
