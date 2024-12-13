package main

import "fmt"

// main is the entry point of the program.
// It demonstrates type conversion between different integer types and string manipulation.
func main() {
	// Declare an int32 variable with a value of 129
	var nilai32 int32 = 129

	// Convert int32 to int64
	var nilai64 int64 = int64(nilai32)

	// Convert int32 to int8, which will cause overflow
	// -128 <= int8 <= 127, so 129 becomes -127
	var nilai8 int8 = int8(nilai32)

	// Print the values of the variables
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// Declare a string variable
	var name = "fathan"

	// Get the first character of the string and convert it to a string
	var f = string(name[0])

	// Print the first character of the string
	fmt.Println(f)
}
