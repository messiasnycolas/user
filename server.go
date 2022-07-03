package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/messiasnycolas/user/handler"
)

var port int = 3000

func main() {
	http.HandleFunc("/user/", handler.UserRoutes)
	log.Printf("Server listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
