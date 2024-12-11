package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "fathan"
	middleName = "abdul"
	lastName = "shodiq"

	return
}

func main() {
	firstName, _, lastName := getCompleteName()

	fmt.Println(firstName, lastName)
}
