package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hallo ini adalah halaman index")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hallo ini adalah halaman utama")
}

func nama(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hallo nama saya adalah Eko Kurniadi")
}

func templates(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"Name":    "Eko Kurniadi",
		"Message": "Saya belajar web golang",
		"Hobi":    "silahkan tulis hobi masing-masing",
	}

	tampil, nilaiError := template.ParseFiles("template.html")
	if nilaiError != nil {
		fmt.Println(nilaiError.Error())
		return
	}

	tampil.Execute(w, data)
}

func belajarHtml(w http.ResponseWriter, r *http.Request) {

	var data = map[string]string{
		"No1":      "1",
		"Nama1":    "Joko",
		"Makanan1": "Mie Ayam",
		"Minuman1": "Es Teh",
		"No2":      "2",
		"Nama2":    "Budi",
		"Makanan2": "Bakso",
		"Minuman2": "Es Jeruk",
		"No3":      "3",
		"Nama3":    "Yudi",
		"Makanan3": "Gado-gado",
		"Minuman3": "Kopi",
	}

	tampil, nilaiError := template.ParseFiles("belajar_dasar_html.html")
	if nilaiError != nil {
		fmt.Println(nilaiError.Error())
	}

	tampil.Execute(w, data)

}

func main() {

	http.HandleFunc("/template", templates)
	http.HandleFunc("/nama", nama)
	http.HandleFunc("/", home)
	http.HandleFunc("/index", index)

	// router belajar dasar HTML
	http.HandleFunc("/belajar_html", belajarHtml)

	fmt.Println("Server website berjalan pada http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
