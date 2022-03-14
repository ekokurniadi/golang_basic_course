package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	tampil, nilaiError := template.ParseFiles("template.html")

	if nilaiError != nil {
		fmt.Println(nilaiError.Error())
		return
	}

	var data = map[string]string{
		"Konten": "Ini Halaman Utama",
	}

	tampil.Execute(w, data)

}

func tentang(w http.ResponseWriter, r *http.Request) {
	tampil, nilaiError := template.ParseFiles("tentang.html")

	if nilaiError != nil {
		fmt.Println(nilaiError.Error())
		return
	}

	var data = map[string]string{
		"Konten": "Ini Halaman Tentang",
	}

	tampil.Execute(w, data)
}

func galery(w http.ResponseWriter, r *http.Request) {
	tampil, nilaiError := template.ParseFiles("galery.html")

	if nilaiError != nil {
		fmt.Println(nilaiError.Error())
		return
	}

	var data = map[string]string{
		"Konten": "Ini Halaman Galery",
	}

	tampil.Execute(w, data)
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/tentang", tentang)
	http.HandleFunc("/galery", galery)

	// memanggil file template.css,tentang.css secara static
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//	/assets/static/template.css -> tanpa stripPrefix
	//	/assets/template.css -> dengan stripPrefix

	fmt.Println("Server dijalankan di http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
