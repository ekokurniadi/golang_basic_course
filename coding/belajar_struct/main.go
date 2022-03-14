package main

import "fmt"

//buat sebuah struct dengan nama People
type People struct {
	Name    string
	Age     int
	Address string
	Hobbies []string
}

func main() {
	//buat 1 variabel bernama eko yang memiliki tipe data struct people
	var eko People
	//isi field name dengan nama Eko Kurniadi
	eko.Name = "Eko Kurniadi"
	//isi field age dengan 25
	eko.Age = 25
	//isi field address dengan Jambi, Indonesia
	eko.Address = "Jambi, Indonesia"
	//isi field hobbies yang bertipe array string
	listOfHobbies := []string{"Menonton", "Membaca buku"}
	eko.Hobbies = listOfHobbies

	fmt.Println("-----------------")
	fmt.Println("Memanggil struct people dengan cara keyword var")
	fmt.Println(eko)
	fmt.Println("-----------------")

	//cara pemanggilan atau pengisian field struct people dengan cara struct literals
	fmt.Println("Cara pemanggilan atau pengisian field struct people dengan cara struct literals")
	hobbiesOfJoko := []string{"Memancing", "Menabung", "Membantu orang tua"}
	joko := People{
		Name:    "Joko Anwar",
		Age:     37,
		Address: "Malang",
		Hobbies: hobbiesOfJoko,
	}
	fmt.Println(joko)
	fmt.Println("-----------------")

	//cara pemanggilan atau pengisian field struct people dengan cara struct literals tanpa menyebutkan
	//nama field nya
	fmt.Println("Cara pemanggilan atau pengisian field struct people dengan cara struct literals tanpa menyebutkan nama field nya")
	newJoko := People{"Joko New", 37, "Malang", hobbiesOfJoko}
	fmt.Println(newJoko)
	fmt.Println("-----------------")

	//mengisi dan memanggil struct method people dari function sayHello
	fmt.Println("Mengisi dan memanggil struct method people dari function sayHello")
	helloPeople := People{
		Name: "Eko Kurniadi",
	}

	helloPeople.sayHello()
	fmt.Println("-----------------")

	helloAge := People{
		Age: 25,
	}
	helloAge.sayAge()

	hobbiSaya := []string{"Makan", "Tidur", "Mobile Legend"}
	helloHobbies := People{
		Hobbies: hobbiSaya,
	}
	helloHobbies.sayHobbies()

}

//membuat struct method
func (people People) sayHello() {
	fmt.Println("Hy, nama saya adalah", people.Name)
}

func (people People) sayAge() {
	fmt.Println("Hy, umur saya adalah", people.Age)
}

func (people People) sayHobbies() {
	fmt.Println("Hy, hobbi saya adalah", people.Hobbies)
}
