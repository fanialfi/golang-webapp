package main

import (
	"log"
	"net/http"
	"time"

	"golang-webapp/handler"
)

func main() {
	// HandleFunc digunakan sebagai routing, parameter pertama adalah route nya, dan kedua adalah handler nya
	// jika mengakses route yang tidak terdaftar, maka secara otomatis handle route / akan terpanggil
	http.HandleFunc("/index", handler.HandleIndex)
	http.HandleFunc("/hello", handler.HandleHello)

	// routing static asset
	http.Handle("/static/",
		http.StripPrefix("/static", // hanya digunakan untuk membungkus actual handler (StripPrefix berguna untuk menghapus prefix dari endpoint yang di request)
			http.FileServer(http.Dir("assets")))) // actual handler

	// render html template
	http.HandleFunc("/", handler.HandleHtmlTemplate)

	var port = ":8080"
	log.Printf("server start at localhost%s\n", port)

	// ListenAndServe digunakan untuk membuat sekaligus memulai server baru,
	// parameter pertama adalah alamat web server, bisa diisi host & port / port saja
	// parameter kedua merupakan object mux / multiplexer (menggunakan default mux yang disediakan oleh golang, jadi diisi dengan nil)
	// err := http.ListenAndServe(port, nil)
	// if err != nil {
	// log.Println("Error :", err.Error())
	// }

	// kelebihan menggunakan http.Server salah satunya adalah kemampuan untuk mengubah konfigurasi default web server go
	var server http.Server
	server.Addr = port
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error :", err.Error())
	}
}
