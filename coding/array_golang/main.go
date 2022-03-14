package main

import "fmt"

func main() {
	var fruits [4]string
	fruits[0] = "Anggur"
	fruits[1] = "Apel"
	fruits[2] = "Mangga"
	fruits[3] = "Jeruk"

	//cara deklarasi array dan cara pengisian elemen-elemen array nya
	fmt.Println("cara deklarasi array dan cara pengisian elemen-elemen array nya")
	fmt.Println("Elemen array pada variabel fruits adalah ", fruits)
	fmt.Println("---------------------------------------------------")

	var buah = [4]string{"Apel", "Semangka", "Jeruk", "Pepaya"}
	//cara deklarasi array dan cara pengisian elemen-elemen array nya
	fmt.Println("cara deklarasi array dan cara pengisian elemen-elemen array nya")
	fmt.Println("Elemen array pada variabel buah adalah ", buah)
	fmt.Println("---------------------------------------------------")

	var arrayBoolean = [5]bool{true, false, false, true, false}
	//cara deklarasi array dan cara pengisian elemen-elemen array nya dengan tipe data array boolean
	fmt.Println("cara deklarasi array dan cara pengisian elemen-elemen array nya dengan tipe data array boolean")
	fmt.Println("Elemen array pada variabel arrayBoolean adalah ", arrayBoolean)
	fmt.Println("---------------------------------------------------")

	var numbers = [...]int{1, 2, 3, 4}
	//cara deklarasi array dan cara pengisian elemen-elemen array nya dengan tipe data int
	fmt.Println("cara deklarasi array dan cara pengisian elemen-elemen array nya dengan tipe data int")
	fmt.Println("Elemen array pada variabel numbers adalah ", numbers)
	fmt.Printf("Jumlah elemen pada variabel numbers adalah : %d\n", len(numbers))
	fmt.Println("---------------------------------------------------")

	var hobbies = [...]string{"Programming", "Membaca buku", "Bermain Sepak Bola", "Belajar Program"}

	//Perulangan variabel array menggunakan for dan len array
	fmt.Println("Perulangan variabel array menggunakan for dan len array")
	fmt.Println("---------------------------------------------------")
	for i := 0; i < len(hobbies); i++ {
		fmt.Printf("Nama saya Eko Kurniadi, hobi saya adalah %s\n", hobbies[i])
	}

	//Perulangan variabel array menggunakan for dan len array
	fmt.Println("Perulangan variabel array menggunakan for dan range dari array")
	fmt.Println("---------------------------------------------------")
	for i, hobbie := range hobbies {
		fmt.Printf("Index array hobbies : %d \t dengan isi elemen : %s\n", i, hobbie)
	}

	//Perulangan variabel array menggunakan for dan len array
	fmt.Println("Perulangan variabel array menggunakan for dan range dari array tanpa menampilkan variabel i")
	fmt.Println("---------------------------------------------------")
	for _, elemen := range hobbies {
		fmt.Printf("Isi elemen Array Hobbies : %s\n", elemen)
	}

	//Soal
	var bilangan = [...]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i := 0; i <= len(bilangan); i++ {
		if i%2 == 0 {
			fmt.Println("Bilangan ganjil ", i)
		}
	}

}
