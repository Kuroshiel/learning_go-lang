package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Eko")
	fmt.Println(result)

	fmt.Println(helper.Application)
	/* fmt.Println(helper.version)
	fmt.Println(helper.sayGoodBye("Eko")) */ // tidak bisa diakses
}
