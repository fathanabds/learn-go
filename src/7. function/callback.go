package main

import (
	"fmt"
)

type Filter func(string) string

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	}
	return name
}

func sayHelloWFilter(name string, filter Filter) string {
	return "hello " + filter(name)
}

func main() {
	filter := func(name string) string {
		if name == "anjing" {
			return "..."
		}
		return name
	}

	fmt.Println(sayHelloWFilter("anjing", spamFilter))
	fmt.Println(sayHelloWFilter("abdul", filter)) // anonymous function
	fmt.Println(sayHelloWFilter("fathan", func(name string) string {
		if name == "anjing" {
			return "..."
		}
		return name
	})) // argument kedua adalah anonymous function
}
