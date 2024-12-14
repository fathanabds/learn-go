package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "insert your host")
	user := flag.String("user", "user", "insert your user")
	password := flag.String("password", "password", "insert your password")

	flag.Parse()

	fmt.Println("host:", *host)
	fmt.Println("user:", *user)
	fmt.Println("password:", *password)
}
