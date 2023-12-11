package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	filteredNamed := filter(name)
// 	fmt.Println("Hello ", filteredNamed)
// }

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredNamed := filter(name)
	fmt.Println("Hello ", filteredNamed)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Eko", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
