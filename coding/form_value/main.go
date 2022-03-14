package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexLogin)
	http.HandleFunc("/login", indexProsesLogin)
	http.HandleFunc("/form", indexForm)
	http.HandleFunc("/proses", indexResult)

	fmt.Println("Server dijalankan pada http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func indexForm(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("form.html"))

		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}

	http.Error(w, "", http.StatusBadRequest)
}

func indexLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form-login").ParseFiles("login.html"))

		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}

	http.Error(w, "", http.StatusBadRequest)
}

func indexResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("result.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var name = r.FormValue("name")
		var message = r.FormValue("message")

		var data = map[string]string{"name": name, "message": message}

		tmpl.Execute(w, data)
	}

	http.Error(w, "", http.StatusBadRequest)
}

func indexProsesLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("form").ParseFiles("form.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tmpl.Execute(w, nil)
	}

	http.Error(w, "", http.StatusBadRequest)
}
