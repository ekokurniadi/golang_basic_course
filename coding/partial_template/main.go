package main

import (
	"fmt"
	"net/http"

	"html/template"
)

type M map[string]interface{}

func main() {
	fmt.Println("Menampilkan partial template")

	var tmpl, err = template.ParseGlob("views/*")

	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Superman"}

		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}

		err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server dijalankan pada localhost:9000")
	http.ListenAndServe(":9000", nil)

}
