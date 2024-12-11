package main

import "fmt"

func sayHello(name string) string {
	return "hello " + name
}

func getFullName(firstName, lastName string) []string {
	return []string{firstName, lastName}
}

func getNewName(firstName, lastName string) (string, string) {
	return firstName, lastName
} // return multiple values

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(sayHello("fathan"))

	fmt.Println(sum(1, 2))

	fmt.Println(getFullName("fathan", "shodiq"))

	firstName, _ := getNewName("eko", "kurniawan") // return multiple values
	fmt.Println(firstName)
}
