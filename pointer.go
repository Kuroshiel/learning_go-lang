package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // copy value
	address2 := &address1

	address2.City = "Bandung"
	fmt.Println(address1) // tidak berubah
	fmt.Println(address2) // berubah menjadi bandung

}
