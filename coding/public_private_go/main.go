package main

import (
	"fmt"
	"public_private_go/library"
	"public_private_go/lingkaran"
	"public_private_go/persegi"
	persegipanjang "public_private_go/persegi_panjang"
)

func main() {
	library.SayHello()
	library.Introduce("Eko Kurniadi")

	fmt.Println("=====Menghitung Persegi=====")
	hitung := persegi.NewHitung(10.0)
	fmt.Println(hitung.Luas())
	fmt.Println(hitung.Keliling())

	fmt.Println("=====Menghitung Persegi Panjang=====")
	hitungPersegiPanjang := persegipanjang.NewHitung(5.0, 10.0)
	fmt.Println(hitungPersegiPanjang.Luas())
	fmt.Println(hitungPersegiPanjang.Keliling())

	fmt.Println("====Menghitung Jari-jari lingkaran")

	hitungLingkaran := lingkaran.NewHitung(50.0)
	fmt.Println(hitungLingkaran.JariJari())
}
