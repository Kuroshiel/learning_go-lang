package main

import "fmt"

func main() {
	var name1 = "Kuro"
	var name2 = "Kuro"

	var result bool = name1 ~= name2

	fmt.Println(result)
}
