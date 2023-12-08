package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello, Eko")

	case "Budi":
		fmt.Println("Hello, Budi")

	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	name = "Khannedy"
	lenght := len(name)
	switch {
	case lenght > 10:
		fmt.Println("Nama Terlalu Panjang")
	case lenght > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
