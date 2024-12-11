package main

import "fmt"

// main is the entry point for the program.
// It demonstrates the use of type declarations in Go.
func main() {
	// NIK is a custom type based on the string type.
	type NIK string

	// NIKFathan is a variable of type NIK, initialized with the value "3216".
	var NIKFathan NIK = "3216"

	// Print the value of NIKFathan to the console.
	fmt.Println(NIKFathan)
}
