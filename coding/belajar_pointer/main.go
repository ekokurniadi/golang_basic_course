package main

import "fmt"

func main() {
	fmt.Println("Coba menampilkan data tanpa POINTER")
	var angkaAwal int = 4
	var angkaB int = angkaAwal

	fmt.Println("Angka Awal :", angkaAwal)
	fmt.Println("Angka B:", angkaB)
	fmt.Println("<------------>")
	fmt.Println("")
	fmt.Println("Coba merubah nilai dari angka awal")
	angkaAwal = 10
	fmt.Println("Angka Awal :", angkaAwal)
	fmt.Println("Alamat dari Angka Awal :", angkaAwal)
	fmt.Println("Angka B:", angkaB)
	fmt.Println("Alamat dari Angka B :", angkaB)
	fmt.Println("<------------>")

	type Address struct {
		Kota, Provinsi, Negara string
	}
	fmt.Println("Menampilkan data address tanpa pointer")
	var address1 Address = Address{
		Kota:     "Jambi",
		Provinsi: "Jambi",
		Negara:   "Indonesia",
	}
	var address2 = address1
	address2.Kota = "Jakarta"
	address2.Provinsi = "DKI Jakarta"
	fmt.Println("Address 1 :", address1)
	fmt.Println("Alamat Address 1,", address1)
	fmt.Println("Address 2 :", address2)
	fmt.Println("Alamat Address 2", address2)
	fmt.Println("")
	fmt.Println("Menampilkan data address dengan pointer")
	var address1Pointer Address = Address{
		Kota:     "Jambi",
		Provinsi: "Jambi",
		Negara:   "Indonesia",
	}
	var address2Pointer *Address = &address1Pointer
	address2Pointer.Kota = "Jakarta"
	address2Pointer.Provinsi = "DKI Jakarta"
	fmt.Println("Address 1 Pointer :", address1Pointer)
	fmt.Println("Alamat Address 1 Pointer", &address1Pointer)
	fmt.Println("Address 2 Pointer :", *address2Pointer)
	fmt.Println("Alamat Address 2 Pointer", address2Pointer)

	//Pertama buat 1 folder project yang baru dengan nama tugas_pointer
	//Buat file go.mod
	//Buat file main.go

	//didalam file main.go buat 1 buah struct dengan nama People yang berisi
	//Name string
	//Age int
	//Address string
	//Buatlah println dan juga variabel kurang lebih sama pada contoh var address1
	//dan var address2 diatas yang tidak menggunakan pointer dan juga contoh yang
	//menggunakan pointer

}
