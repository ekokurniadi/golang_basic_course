package main

import "fmt"

func main() {
	//Seleksi Kondisi menggunakan if
	var point = 8

	if point == 7 {
		fmt.Println("Kondisi If Nilai Point adalah ", point)
	} else {
		fmt.Println("Kondisi Else Nilai Point adalah ", point)
	}

	//Seleksi kondisi menggunakan if dan else if
	if point == 5 {
		fmt.Println("Kondisi pertama if Nilai Point adalah ", point)
	} else if point == 6 {
		fmt.Println("Kondisi kedua else if Nilai Point adalah ", point)
	} else if point == 7 {
		fmt.Println("Kondisi ketiga else if Nilai Point adalah ", point)
	} else if point == 9 {
		fmt.Println("Kondisi keempat else if Nilai Point adalah ", point)
	} else {
		fmt.Println("Kondisi kelima else Nilai Point adalah ", point)
	}

	//Seleksi kondisi menggunakan switch case
	nilai := 78

	switch nilai {
	case 80:
		fmt.Printf("Selamat anda lulus dengan Grade B dan nilai %d\n", 80)
	case 50:
		fmt.Printf("Maaf anda belum lulus dengan Grade C dan nilai %d\n", 50)
	default:
		fmt.Printf("Selamat anda lulus dengan sempurna dengan Grade A dan nilai %d\n", nilai)
	}

	switch {
	case (nilai >= 65) && (nilai <= 70):
		fmt.Println("Grade C")
	case (nilai > 71) && (nilai <= 80):
		fmt.Println("Grade B")
	case (nilai > 81) && (nilai <= 90):
		fmt.Println("Grade B+")
	case (nilai > 91) && (nilai <= 100):
		{
			fmt.Println("Keren")
			fmt.Println("Grade A")
		}
	default:
		fmt.Println("Grade D")
	}

	//buat 1 variabel bernama angka
	//buatlah seleksi kondisi menggunakan switch case
	//yang mana jika case nya 1 maka terminal akan melakukan print perkalian 1
	//menggunakan perulangan for

	angka := 1

	switch angka {
	case 1:
		for i := 1; i <= 10; i++ {
			fmt.Printf("Perkalian %d * %d = %d\n", angka, i, angka*i)
		}
	default:
		fmt.Println("Perkalian tidak ditemukan")
	}
}
