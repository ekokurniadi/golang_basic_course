package main

import (
	"fmt"
)

func main() {

	var firstName string  //ini untuk deklarasi variabel firstName dengan tipe string
	var lastName string   //ini untuk deklarasi variabel lastName dengan tipe string
	var age int           //ini untuk deklarasi variabel age (umur) dengan tipe data integer / angka
	firstName = "Eko"     //disini saya akan isi variabel firtsname dengan nama eko
	lastName = "Kurniadi" //disini saya akan isi variabel lastname dengan nama kurniadi
	age = 25              //disini saya akan isi variabel age dengan angka 25

	fmt.Printf("Hallo %s %s umur %d tahun\n", firstName, lastName, age)
	/*fmt.printf("Hallo akan menulis tulisan Hallo
	, kemudian tanda %s pertama untuk menampung nilai variabel firstName,
	, kemudian tanda %s kedua untuk menampung nilai variabel lastName,
	, %d akan menampung nilai dari variabel age atau umur dengan nilai angka
	")*/

	namaDepan := "Eko"
	namaBelakang := "Kurniadi"
	umur := 25

	fmt.Printf("Hallo %s %s umur %d tahun\n", namaDepan, namaBelakang, umur)

	firts, second, third := "satu", "dua", 3

	fmt.Printf("%s %s %d", firts, second, third)

	for i := 0; i < 10; i++ {
		fmt.Printf("Hallo ke %d %s %s umur %d tahun\n", i, namaDepan, namaBelakang, umur)
	}
}
