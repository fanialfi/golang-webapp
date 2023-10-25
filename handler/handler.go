package handler

import (
	"html/template"
	"net/http"
	"path"
)

func HandleIndex(res http.ResponseWriter, req *http.Request) {
	message := "Welcome"
	res.Write([]byte(message))
}

func HandleHello(res http.ResponseWriter, req *http.Request) {
	message := "Hello World"
	res.Write([]byte(message))
}

// handler for render html template
func HandleHtmlTemplate(res http.ResponseWriter, req *http.Request) {
	var filePath = path.Join("views", "index.html")

	// template.ParseFiles digunakan untuk parsing file template
	var tmpl, err = template.ParseFiles(filePath)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]any{
		"title": "Learning Golang Web",
		"name":  "Fani",
	}

	// digunakan untuk menyisipkan data pada template, kemudian ditampilkan ke parameter pertama (res / browser)
	// data yang disisipkan dapat berupa struct, map, atau interface kosong
	err = tmpl.Execute(res, data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
