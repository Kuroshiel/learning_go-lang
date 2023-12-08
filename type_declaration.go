package main

import "fmt"

func main() {

	type NoKTP string

	var ktpEko NoKTP = "11111111"

	var contoh string = "22222222"
	var contohktp NoKTP = NoKTP(contoh)

	fmt.Println(ktpEko)
	fmt.Println(contohktp)
}
