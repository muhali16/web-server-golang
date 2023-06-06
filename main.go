package main

import (
	. "fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/mhs", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]interface{}{
			"nama": "Muhammad Ali Mustaqim",
			"npm":  202110715138,
		}
		var tmplt, e = template.ParseFiles("mhs.html")
		if e != nil {
			Println(e.Error())
			return
		}
		tmplt.Execute(w, data)
	})
	http.HandleFunc("/intostruct", func(w http.ResponseWriter, r *http.Request) {
		Fprintln(w, "<h1>Decode Json Strings Into Struct</h1>")
		var jsonStrings = `{"nama": "Laurence NS", "usia": 21}`
		var jsonData = decodeJsonIntoStruct(jsonStrings)
		Fprintln(w, "Nama:", jsonData.Nama, "<br>")
		Fprintln(w, "Usia:", jsonData.Usia)
	})
	http.HandleFunc("/intomap", func(w http.ResponseWriter, r *http.Request) {
		Fprintln(w, "<h1>Decode Json Strings Into Map</h1>")
		var jsonStrings = `{"nama": "Ali", "usia": 20}`
		var jsonData = decodeJsonIntoMap(jsonStrings)
		Fprintln(w, "Nama:", jsonData["nama"], "<br>")
		Fprintln(w, "Usia:", jsonData["usia"])
	})
	http.HandleFunc("/arrayjson", func(w http.ResponseWriter, r *http.Request) {
		Fprintln(w, "<h1>Decode Array Json Strings Into Array Object</h1>")
		var jsonString = `[
			{"nama": "Ali", "usia": 20},
			{"nama": "Laurence NS", "usia": 21}
		]`
		var jsonData = decodeArrayJsonIntoArrayObject(jsonString)
		Println(jsonData)
		Fprintln(w, "<ol>")
		for _, m := range jsonData {
			Fprintln(w, "<li>Nama:", m.Nama, "| Usia:", m.Usia, "</li>")
		}
		// Fprintln(w, "<li>Nama:", jsonData[0].Nama, "| Usia:", jsonData[0].Usia, "</li>")
		// Fprintln(w, "<li>Nama:", jsonData[1].Nama, "| Usia:", jsonData[1].Usia, "</li>")
		Fprintln(w, "</ol>")
	})
	http.HandleFunc("/intojson", func(w http.ResponseWriter, r *http.Request) {
		var user = []User{
			{"Ali", 20},
			{"Laurence", 21},
			{"Budhi", 22},
			{"Adhit", 23},
		}
		var jsonData = encodeObjectIntoJson(&user)

		Fprintln(w, "<h1>Encode Array Object Into Array Json</h1>")
		Fprintln(w, jsonData)
	})
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		var students, err = fetchStudents()
		if err != nil {
			Println(err.Error())
			return
		}
		Fprintln(w, "<h1>HTTP Client Golang</h1>")
		Fprintln(w, "<ol>")
		for _, m := range students {
			Fprintln(w, "<li>NIS:", m.Id, "| Nama:", m.Name, "| Nilai:", m.Grade, "</li>")
		}
		Fprintln(w, "</ol>")
	})
	http.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		Fprintln(w, "<h1>HTTP Client Golang 2</h1>")
		var student, err = fetchStudent(Sprint(r.FormValue("id")))
		if err != nil {
			Fprintln(w, "<h2>Data Not Found - 404</h2>")
		}
		Fprintln(w, "<ul>")
		Fprintln(w, "<li>Nama: ", student.Name, "</li>")
		Fprintln(w, "<li>NIS: ", student.Id, "</li>")
		Fprintln(w, "<li>Nilai: ", student.Grade, "</li>")
		Fprintln(w, "</ul>")
	})

	http.HandleFunc("/carimhs", func(w http.ResponseWriter, r *http.Request) {
		var tmplte, err = template.ParseFiles("searchmhs.html")
		if err != nil {
			Println(err.Error())
			return
		}

		tmplte.Execute(w, nil)
	})

	Println("Server running at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
