package main

import "fmt"

func main() {
	var a = 10;
	var b = 10;
	var c = a + b
	fmt.Println(c)

	a += 5 // a = a + 5
	fmt.Println(a)

	b++; // b = b + 1
	fmt.Println(b)
}
