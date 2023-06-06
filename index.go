package main

import (
	. "fmt"
	"net/http"
	"net/url"
)

func index(w http.ResponseWriter, r *http.Request) {
	Fprintln(w, "<h1>Selamat datang, Fellas!</h1>")
	Fprintln(w, "<hr>")
	Fprintln(w, "<h2>URL Parse Golang</h2>")
	Fprintln(w, "<p>Contoh URL: http://localhost:8080/kelas?id=2&delete=false</p>")
	var u, err = url.Parse("http://localhost:8080/kelas?id=2&delete=false")
	if err != nil {
		Println(err.Error())
		return
	}
	Fprintln(w, "Host:"+u.Host+"<br>")
	Fprintln(w, "Path:"+u.Path+"<br>")
	Fprintln(w, "Id:"+u.Query().Encode()+"<br>")
}
