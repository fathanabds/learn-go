package main

import "fmt"

// The main package is required for executable programs
// Importing the fmt package for formatted I/O
func main() {
	person := map[string]string{
		"name": "fathan",
		"city": "bekasi",
	}

	person["title"] = "student"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["title"])

	fmt.Println(len(person))
	delete(person, "city")
	fmt.Println(person)
}
