package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Address     string
	Affiliation string
}

func (t Info) GetAffiliationDetailInfo() string {
	return `Have 31 divisions`
}

type Person struct {
	Info   Info
	Name   string
	Gender string
	Hobies []string
}

var port = ":6060"

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		person := Person{
			Name:   "Bruce Wayne",
			Gender: "male",
			Hobies: []string{"Reading Books", "Travelling", "Buying things"},
			Info:   Info{Address: "Gotham City", Affiliation: "Wayne Enterprice"},
		}

		tmpl := template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(res, person); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Printf("Server running on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
