package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Belajar membuat fungsi pada golang tanpa nilai balik")
	var siswa = []string{"Ruby", "Ryu", "Richel"}
	printMessage("Hello", siswa)
	fmt.Println("----------------------------------------------------")

	fmt.Println("Belajar membuat fungsi dengan nilai balik")
	hasilPenjumlahan := hitungAngka(17, 25)
	fmt.Printf("Hasil penjumlahan angka di atas adalah %d\n", hasilPenjumlahan)
	hasilRataRata := rataRata(hasilPenjumlahan, 12)
	fmt.Printf("Hasil Rata-rata angka di atas adalah %d\n", hasilRataRata)
	fmt.Println("----------------------------------------------------")

	fmt.Println("Belajar memanggil fungsi dari lain file")
	PrintHello()
	fmt.Println("Belajar memanggil fungsi dari lain file")
	PrintNama("Eko Kurniadi")
}

// ini adalah fungsi tanpa nilai balik
func printMessage(pesan string, array []string) {
	var gabunganKata = strings.Join(array, " ")
	fmt.Println(pesan, gabunganKata)
}

//ini adalah fungsi yang memiliki nilai balik
func hitungAngka(nilai1, nilai2 int) int {
	//buat 1 variabel dengan nama hasil yang akan menyimpan penjumlahan dari nilai1 + nilai2
	var hasil = nilai1 + nilai2
	//export hasil dari penjumlahan tersebut agar bisa digunakan pada fungsi lainnya
	return hasil
}
func rataRata(nilai1, nilai2 int) int {
	var hasilRataRata = nilai1 / nilai2
	return hasilRataRata
}
