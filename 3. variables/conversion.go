package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) // -128 <= int8 <= 127, sehingga 129 jadi -127

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "fathan";
	var f = string(name[0])
	fmt.Println(f)

}