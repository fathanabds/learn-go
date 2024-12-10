package main

import "fmt"

func main(){
	const firstName string = "fathan" // menuliskan tipe data
	fmt.Println(firstName)

	const gpa = 3.37 // tanpa tipe data
	fmt.Println(gpa)

	const (
		lastName = "abdul"
		age = 25
	) // deklarasi multiple constants
}