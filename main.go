package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fileLocation := filepath.Join("views", "index.html")
	tmplt, err := template.ParseFiles(fileLocation)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	data := map[string]string{
		"title": "Welcome",
		"name":  "Ali",
		"css":   "/static/style.css",
	}

	tmplt.Execute(w, data)
}

func main() {
	// web route
	http.HandleFunc("/", index)
	http.HandleFunc("/welcome", welcome)

	// render static file from directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	address := "localhost:8000"
	fmt.Println("Server running on", address)

	// first way to serve golang web server
	// http.ListenAndServe(address, nil)

	// second way to serve golang web server
	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
	}
}
