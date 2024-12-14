package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "FaTHaN AbDUl"
	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.Contains(strings.ToLower(name), "fa"))

	splittedName := strings.Split(name, " ")
	fmt.Println(splittedName)
}
