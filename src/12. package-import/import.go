package main

import (
	"fmt"
	_ "learn-go/database" // blank identifier
	"learn-go/helpers"
)

func main() {
	fmt.Println(helpers.SayHello("eko"))
	// fmt.Println(database.GetDB())
}
