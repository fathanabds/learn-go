package main

import "fmt"

// main is the entry point of the program.
func main() {
	// Declare a float32 variable named gpa.
	var gpa float32 // deklarasi di awal
	// Assign a value to gpa.
	gpa = 3.37
	// Print the value of gpa.
	fmt.Println(gpa)

	// Declare and initialize a string variable named name.
	var name string = "shodiq" // tipe data diberikan
	// Print the value of name.
	fmt.Println(name)

	// Declare and initialize an integer variable named age without specifying the type.
	var age = 10 // tipe data tidak diberikan, go akan deteksi auto
	// Print the value of age.
	fmt.Println(age)

	// Declare and initialize a boolean variable named isMarried using shorthand notation.
	isMarried := true // tanpa menggunakan keyword 'var'
	// Print the value of isMarried.
	fmt.Println(isMarried)

	// Declare multiple variables firstName and lastName.
	var (
		firstName = "fathan"
		lastName  = "abdul"
	) // deklarasi multiple variable

	// Print the value of firstName.
	fmt.Println(firstName)
	// Print the value of lastName.
	fmt.Println(lastName)
}
