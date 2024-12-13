package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	eko := Man{Name: "Eko"}
	eko.married()
	fmt.Println(eko)
}
