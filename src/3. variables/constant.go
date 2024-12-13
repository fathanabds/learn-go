package main

import "fmt"

// main is the entry point of the program.
func main() {
	// firstName is a constant string with the value "fathan".
	const firstName string = "fathan" // menuliskan tipe data
	fmt.Println(firstName)

	// gpa is a constant float64 with the value 3.37.
	const gpa = 3.37 // tanpa tipe data
	fmt.Println(gpa)

	const (
		// lastName is a constant string with the value "abdul".
		lastName = "abdul"
		// age is a constant integer with the value 25.
		age = 25
	) // deklarasi multiple constants
}
