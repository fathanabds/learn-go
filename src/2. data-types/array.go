package main

import "fmt"

// Import the fmt package for formatted I/O operations.

// main is the entry point of the program.
// It demonstrates the usage of arrays and slices in Go.
func main() {
	// Declare an array of strings with a fixed size of 3 and assign a value to the first element.
	var names [3]string
	names[0] = "fathan"

	// Print the first element of the names array.
	fmt.Println(names[0])

	// Declare and initialize an array of integers with five elements.
	var values = [...]int{
		1, 2, 3, 4, 5,
	}

	// Print the entire values array.
	fmt.Println(values)
	// Print the length of the values array.
	fmt.Println(len(values))

	// Modify the third element of the values array.
	values[2] = 4
	// Print the modified values array.
	fmt.Println(values)
}
