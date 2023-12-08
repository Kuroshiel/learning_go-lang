package main

import "fmt"

func main() {
	// 	var person map[string]string{}
	// 	person["name"] = "Eko"
	// 	person["addres"] = "Subang"

	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Eko Kurniawan"
	book["wrong"] = "Ups"
	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}
