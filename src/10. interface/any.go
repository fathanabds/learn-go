package main

import "fmt"

func ups(i int) any {
	switch i {
	case 1:
		return 1
	case 2:
		return true
	default:
		return "ups"
	}
}

func main() {
	fmt.Println(ups(1))
	fmt.Println(ups(2))
	fmt.Println(ups(3))
}
