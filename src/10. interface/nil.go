package main

import (
	"fmt"
)

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func checkNil(data map[string]string) {
	if data == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(data)
	}
}

func main() {
	person := newMap("Fathan")
	checkNil(person)
	person2 := newMap("")
	checkNil(person2)
}
