package main

import (
	"fmt"
	"strings"
)

func main() {
	test := "Here is $1 which is $2"

	// replace any instance of $1 with honey and $2 with tastey
	test = strings.ReplaceAll(test, "$1", "honey")
	test = strings.ReplaceAll(test, "$2", "tastey")

	fmt.Println(test)
}