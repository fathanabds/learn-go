package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, error := strconv.ParseBool("false")
	if error == nil {
		fmt.Println(boolean)
	} else {
		fmt.Print(error.Error())
	}

	number, error := strconv.ParseInt("1", 10, 32)
	if error == nil {
		fmt.Println(number)
	} else {
		fmt.Print(error.Error())
	}

	fmt.Println(strconv.FormatInt(1000, 10))
}
