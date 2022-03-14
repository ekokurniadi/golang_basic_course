package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Belajar Golang HTTP Method")

	http.HandleFunc("/post", func(tulis http.ResponseWriter, baca *http.Request) {
		method := baca.Method

		text := ""
		text = "Hello ini adalah http method " + method

		tulis.Write([]byte(text))
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method

		text := ""
		text = "Hello ini adalah http method" + method

		w.Write([]byte(text))
	})

	http.HandleFunc("/put", func(rw http.ResponseWriter, r *http.Request) {
		method := r.Method

		text := ""
		text = "Hello ini adalah method " + method

		rw.Write([]byte(text))
	})

	fmt.Println("Server dijalankan pada localhost:9000")
	fmt.Println("Tuliskan perintah http://localhost:9000/post untuk menjalankan method post")
	fmt.Println("Tuliskan perintah http://localhost:9000/get untuk menjalankan method get")

	http.ListenAndServe(":9000", nil)
}
