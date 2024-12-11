package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("selesai melaksanakan function")
	message := recover()
	if message != nil {
		fmt.Println("terjadi error", message)
	}
}

func runApplication(error bool) {
	defer endApp() // mirip finally di js, fn ini akan selalu dijalankan walaupun terjadi error di fn runApp
	defer func() {
		fmt.Println("ini anonim fn")
	}() // defer menggunakan anonim fn
	if error {
		panic("APLIKASI ERROR") // mirip throw di js
	}

	fmt.Println("run app")
}

func main() {
	runApplication(true)
	fmt.Println("fathan")
}
