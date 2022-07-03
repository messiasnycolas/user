package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/messiasnycolas/user/config"
	handler "github.com/messiasnycolas/user/handler"
)

func main() {
	config.Load()

	var port string = config.GetAPIConfig().Port
	http.HandleFunc("/user/", handler.UserRoutes)

	log.Printf("Server listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
