package main

import "fmt"

func main() {
	name := "abdul"

	if name == "fathan" {
		fmt.Println("hello", name)
	} else if name == "abdul" {
		fmt.Println("welcome back", name)
	} else {
		fmt.Println("who r u")
	}

	city := "bekasi"
	if len(city) > 5 {
		fmt.Println("you are here!")
	}
}
