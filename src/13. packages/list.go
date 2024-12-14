package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Fathan")
	data.PushBack(1)
	data.PushBack(true)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	} // dari depan => belakang

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	} // dari belakang => depan
}
