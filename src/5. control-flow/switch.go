package main

import "fmt"

func main() {
	name := "shodiq"

	switch name {
	case "abdul":
		fmt.Println("hello", name)
		fmt.Println("how r u")
	case "fathan":
		fmt.Println("welcome back", name)
	default:
		fmt.Println("who r u")
	}

	length := len(name)

	switch {
	case length <= 5:
		fmt.Println("nama sudah benar")
	default:
		fmt.Println("nama terlalu panjang")
	}
}
