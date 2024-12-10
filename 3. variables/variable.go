package main

import "fmt";

func main(){
	var gpa float32; // deklarasi di awal
	gpa = 3.37
	fmt.Println(gpa)

	var name string = "shodiq"; // tipe data diberikan
	fmt.Println(name)

	var age = 10; // tipe data tidak diberikan, go akan deteksi auto
	fmt.Println(age)

	isMarried := true // tanpa menggunakan keyword 'var'
	fmt.Println(isMarried)

	var (
		firstName = "fathan"
		lastName = "abdul"
	) // deklarasi multiple variable

	fmt.Println(firstName)
	fmt.Println(lastName)
}