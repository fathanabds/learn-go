package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println("")
	names := []string{"fathan", "abdul", "shodiq"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, name := range names {
		fmt.Println("index", i, "=", name)
		// fmt.Println(name)
	}

	person := map[string]string{
		"name":  "fathan",
		"title": "student",
	}

	for key, value := range person {
		fmt.Println("key", key, "=", value)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Print(i, " ")
	}

	fmt.Println("")
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
}
