package main

import "fmt"

// main is the entry point of the program.
// It demonstrates the usage of arrays and slices in Go.
func main() {
	// Declare and initialize an array of integers with 5 elements.
	var values = [...]int{
		1, 2, 3, 4, 5,
	}

	// Create a slice from the entire array using the append function.
	var slice1 = append(values[:])
	// Modify the first element of the slice.
	slice1[0] = 10
	// Print the modified slice.
	fmt.Println(slice1)
	// Print the original array to show it remains unchanged.
	fmt.Println(values)

	// Create a slice from the values array starting from index 1 up to (but not including) index 3.
	var sliceExample = values[1:3]
	// Print the created slice.
	fmt.Println(sliceExample)
	// Print the length of the created slice.
	fmt.Println(len(sliceExample))
	// Print the capacity of the created slice.
	fmt.Println(cap(sliceExample))
}
