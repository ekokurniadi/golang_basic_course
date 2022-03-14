package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	var address1 Address = Address{
		City:     "Jambi",
		Province: "Provinsi Jambi",
		Country:  "Indonesia",
	}
	var address2 *Address = &address1
	address2.City = "Jakarta"

	address2 = &Address{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
}
