package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var port = ":8080"

type SuperHero struct {
	Name    string
	Alias   string
	Friends []string
}

func (s SuperHero) SayHello(from, message string) string {
	return fmt.Sprintf(`%s said : "%s"`, from, message)
}

func handleRoot(res http.ResponseWriter, req *http.Request) {
	var person = SuperHero{
		Name:    "Bruce Wayne",
		Alias:   "Batman",
		Friends: []string{"Superman", "Flash", "Green Lantern"},
	}

	var tmpl = template.Must(template.ParseFiles("view.html"))
	if err := tmpl.Execute(res, person); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handleRoot)

	log.Printf("server running on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
