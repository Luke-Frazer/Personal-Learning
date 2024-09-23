package main

import "fmt"

func main() {
	a, b := 12, 345
	c, d := 1.2, 3.45

	// examples of how to print different numbers, and format into a clean table
	fmt.Printf("|%9d|%9d|\n", a, b)
	fmt.Printf("|%09d|%09d|\n", a, b)
	fmt.Printf("|%-9d|%-9d|\n", a, b)
	fmt.Printf("|%9f|%9.2f|\n", c, d)

}
