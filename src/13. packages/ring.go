package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5) // data => elemen ring yang pertama
	for i := 0; i < data.Len(); i++ {
		data.Value = "value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	// for i := 0; i < data.Len(); i++ {
	// 	fmt.Println(data.Value)
	// 	data = data.Next()
	// }

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
