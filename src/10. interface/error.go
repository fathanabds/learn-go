package main

import (
	"errors"
	"fmt"
)

func divide(num, divider int) (int, error) {
	if divider == 0 {
		return num, errors.New("cannot be divided by 0")
	} else {
		return num / divider, nil
	}
}

func main() {
	var contohError error = errors.New("New Error")
	fmt.Println(contohError.Error())
	result, err := divide(4, 2)
	if err == nil {
		fmt.Println("result:", result)
	} else {
		fmt.Println("error", err.Error())
	}
}
