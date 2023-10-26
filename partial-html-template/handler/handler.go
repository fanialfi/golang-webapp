package handler

import (
	"html/template"
	"net/http"
)

type M map[string]any

func HandleIndex(res http.ResponseWriter, req *http.Request) {
	data := M{"nama": "Batman"}

	// template.Must digunakan untuk deteksi error pada saat pembuatan instance *template.Template
	// ketika ada error, panic akan dimunculkan
	tmpl := template.Must(template.ParseFiles(
		"views/index.html",
		"views/_header.html",
		"views/_message.html",
	))

	err := tmpl.ExecuteTemplate(res, "index", data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func HandleAbout(res http.ResponseWriter, req *http.Request) {
	data := M{"nama": "Batman"}
	tmpl := template.Must(template.ParseFiles(
		"views/about.html",
		"views/_header.html",
		"views/_message.html",
	))

	err := tmpl.ExecuteTemplate(res, "about", data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
