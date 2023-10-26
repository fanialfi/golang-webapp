package main

import (
	"log"
	"net/http"
	"partial-html-template/handler"
)

var port = ":8080"

func main() {
	http.HandleFunc("/index", handler.HandleIndex)
	http.HandleFunc("/about", handler.HandleAbout)

	log.Printf("server running on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
