package main

import "fmt"

func main() {
	const (
		firstName string = "Kurogane"
		lastName         = "Shiro"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Error Jika Diibuah Value/nilainya
	// firstName = "Kuroro"
	// lastName = "Kuroshiel"
}
