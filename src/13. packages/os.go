package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	command := args[1]
	fmt.Println(args)
	fmt.Println(command)

	fmt.Println("====")

	name, err := os.Hostname()
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err.Error())
	}

}
