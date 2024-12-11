package main

import "fmt"

func logging() {
	fmt.Println("selesai melaksanakan function")
}

func runApp(value int) {
	defer logging() // mirip finally di js, fn ini akan selalu dijalankan walaupun terjadi error di fn runApp
	fmt.Println("run app")
	fmt.Println(10 / value)
}

func main() {
	runApp(0)
}
