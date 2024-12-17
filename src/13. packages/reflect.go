package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	name string `required:"true" max:"10"`
	age  int
}

func IsValid(data any) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			return value != ""
		}
	}
	return true
}

func main() {
	sample := Sample{name: "Fathan", age: 25}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.Field(1).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(IsValid(sample))
}
