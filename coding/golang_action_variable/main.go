package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Pekerjaan string
	Alamat    string
}

type Profil struct {
	Nama         string
	JenisKelamin string
	Hobi         []string
	Biodata      Info
}

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		var profil = Profil{
			Nama:         "Ahmad",
			JenisKelamin: "laki-laki",
			Hobi:         []string{"Menonton TV", "Jalan-jalan", "Memancing"},
			Biodata:      Info{Pekerjaan: "Programmer", Alamat: "Jambi"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))

		if err := tmpl.Execute(w, profil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("Server dijalankan pada localhost:9000")
	http.ListenAndServe(":9000", nil)
}
