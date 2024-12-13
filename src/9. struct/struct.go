package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func main() {
	fathan := Customer{Name: "Fathan", Address: "Bekasi", Age: 25} // struct literal
	fmt.Println(fathan)
	fathan.Age = 30
	fmt.Println(fathan)
	fmt.Println(fathan.Age)

	budi := Customer{"Budi", "Depok", 30}
	fmt.Println(budi)

	fathan.sayHello()
}
