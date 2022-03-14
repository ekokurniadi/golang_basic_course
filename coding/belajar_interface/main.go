package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

type persegi struct {
	sisi float64
}

// fungsi menghitung jari-jari lingkaran
func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

//fungsi menghitung luas lingkaran
func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

//fungsi menghitung keliling lingkaran
func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

//fungsi menghitung luas persegi
func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

//fungsi menghitung keliling persegi
func (p persegi) keliling() float64 {
	return p.sisi * 4
}

type persegiPanjang struct {
	panjang float64
	lebar   float64
}

//fungsi menghitung luas persegi panjang
func (nilai persegiPanjang) luas() float64 {
	return nilai.panjang * nilai.lebar
}

func (nilai persegiPanjang) keliling() float64 {
	return (2 * nilai.panjang) + (2 * nilai.lebar)
}

func main() {
	var bangunDatar hitung

	bangunDatar = lingkaran{diameter: 20.0}
	fmt.Println("=========Lingkaran=======")
	fmt.Println("nilai jari-jari :", bangunDatar.(lingkaran).jariJari())
	fmt.Println("nilai luas :", bangunDatar.luas())
	fmt.Println("nilai keliling :", bangunDatar.keliling())

	bangunDatar = persegi{sisi: 15.0}
	fmt.Println("======Persegi======")
	fmt.Println("nilai luas persegi :", bangunDatar.luas())
	fmt.Println("nilai keliling persegi :", bangunDatar.keliling())

	bangunDatar = persegiPanjang{panjang: 20.0, lebar: 5.0}
	fmt.Println("======Persegi Panjang======")
	fmt.Println("luas persegi panjang :", bangunDatar.luas())
	fmt.Println("keliling persegi panjang :", bangunDatar.keliling())
}
