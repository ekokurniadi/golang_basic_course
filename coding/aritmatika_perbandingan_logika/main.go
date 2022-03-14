package main

import "fmt"

func main() {

	nilai1 := 15
	nilai2 := 10

	fmt.Println("Operator Aritmatika")
	// Operator Aritmatika
	hasilPenjumlahan := nilai1 + nilai2
	fmt.Printf("Hasil penjumlahan adalah %d\n", hasilPenjumlahan)

	hasilPengurangan := nilai1 - nilai2
	fmt.Printf("Hasil pengurangan adalah %d\n", hasilPengurangan)

	hasilPerkalian := nilai1 * nilai2
	fmt.Printf("Hasil perkalian adalah %d\n", hasilPerkalian)

	var hasilPembagian float64 = float64(nilai1) / float64(nilai2)
	fmt.Printf("Hasil pembagian adalah %v\n", hasilPembagian)

	hasilModulus := nilai1 % nilai2
	fmt.Printf("Hasil modulus adalah %d\n", hasilModulus)

	fmt.Println("Operator Perbandingan")
	//Operator Perbandingan
	hasilPerbandinganSamaDengan := nilai1 == nilai2
	// Apakah 15 itu sama 10 ?
	fmt.Printf("Hasil perbandingan sama dengan %t\n", hasilPerbandinganSamaDengan)

	hasilPerbandinganTidakSamaDengan := nilai1 != nilai2
	// Apakah 15 bukan 10 ?
	fmt.Printf("Hasil perbandingan tidak sama dengan %t\n", hasilPerbandinganTidakSamaDengan)

	hasilPerbandinganLebihKecilDari := nilai1 < nilai2
	// Apakah 15 lebih kecil dari 10 ?
	fmt.Printf("Hasil perbandingan lebih kecil dari %t\n", hasilPerbandinganLebihKecilDari)

	hasilPerbandinganLebihKecilAtauSamaDengan := nilai1 <= nilai2
	// Apakah 15 lebih kecil atau sama dengan 10 ?
	fmt.Printf("Hasil perbandingan lebih kecil atau sama dengan %t\n", hasilPerbandinganLebihKecilAtauSamaDengan)

	hasilPerbandinganLebihBesarDari := nilai1 > nilai2
	// Apakah 15 lebih besar dari 10 ?
	fmt.Printf("Hasil perbandingan lebih besar dari %t\n", hasilPerbandinganLebihBesarDari)

	hasilPerbandinganLebihBesarAtauSamaDengan := nilai1 >= nilai2
	// Apakah 15 lebih besar atau sama dengan 10 ?
	fmt.Printf("Hasil perbandingan lebih besar atau sama dengan %t\n", hasilPerbandinganLebihBesarAtauSamaDengan)

	fmt.Println("Operator Logika")
	// Operator Logika

	benar := true
	salah := false

	hasilOperatorLogikaDan := benar && salah
	// Apakah benar dan salah ?
	fmt.Printf("Hasil operator logika && = %t\n", hasilOperatorLogikaDan)

	hasilOperatorLogikaAtau := benar || salah
	// Apakah benar atau salah ?
	fmt.Printf("Hasil operator logika || = %t\n", hasilOperatorLogikaAtau)

	hasilOperatorNegasi := !salah
	// Apakah lawan dari salah ?
	fmt.Printf("Hasil operator negasi = %t\n", hasilOperatorNegasi)

	// buatlah perkalian angka 10
	for i := 1; i <= 10; i++ {
		fmt.Printf("Perkalian %d = %d\n", i, i*10)
	}

}
