package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnggunPermata/Simple-Web-with-Golang/handlers"
)

//Setting port
const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println(fmt.Sprintf("Starting on port %s", port))
	_ = http.ListenAndServe(port, nil)
}
