package main

import (
	"log"
	"net/http"
)

func handleIndex(res http.ResponseWriter, req *http.Request) {
	message := "Welcome"
	res.Write([]byte(message))
}

func handleHello(res http.ResponseWriter, req *http.Request) {
	message := "Hello World"
	res.Write([]byte(message))
}

func main() {
	// HandleFunc digunakan sebagai routing, parameter pertama adalah route nya, dan kedua adalah handler nya
	// jika mengakses route yang tidak terdaftar, maka secara otomatis handle route / akan terpanggil
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/index", handleIndex)
	http.HandleFunc("/hello", handleHello)

	var port = ":8080"
	log.Printf("server start at localhost%s\n", port)

	// ListenAndServe digunakan untuk membuat sekaligus memulai server baru,
	// parameter pertama adalah alamat web server, bisa diisi host & port / port saja
	// parameter kedua merupakan object mux / multiplexer (menggunakan default mux yang disediakan oleh golang, jadi diisi dengan nil)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("Error :", err.Error())
	}
}
