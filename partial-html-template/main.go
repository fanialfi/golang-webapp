package main

import (
	"html/template"
	"log"
	"net/http"
)

type M map[string]any

var port = ":8080"

func main() {
	Tmpl, err := template.ParseGlob("./views/*")
	if err != nil {
		panic(err.Error())
		// return
	}

	http.HandleFunc("/index", func(res http.ResponseWriter, req *http.Request) {
		data := M{"nama": "Batman"}
		err = Tmpl.ExecuteTemplate(res, "index", data)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(res http.ResponseWriter, req *http.Request) {
		data := M{"nama": "Batman"}
		err = Tmpl.ExecuteTemplate(res, "about", data)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Printf("server running on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
