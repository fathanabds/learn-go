package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))

	search := regex.FindAllString("eko edo eka emo elu", -1) // -1 berarti semua string yang match
	fmt.Println(search)
}
