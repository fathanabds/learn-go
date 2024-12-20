package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3))

	numbers := []int{10, 10, 10}
	fmt.Println(sumAll(numbers...))
}
