package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplacation() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	runApplacation()
}
