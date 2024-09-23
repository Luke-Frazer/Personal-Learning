package main

import (
	"fmt"
)

func main() {
	a := 2
	b := 3.1

	// %T prints type, %v prints any variable type,
	// the %<number>T causes the value in T to take up <number> amount of space, good for aligning
	// the %[<number>]v allows %v to reuse its value twice
	fmt.Printf("a: %8T %[1]v\n", a)
	fmt.Printf("b: %8T %[1]v\n", b)

	// cannot assign a = b since they are different types
	// However, we can type cast
	a = int(b)
	fmt.Printf("a: %8T %[1]v\n", a)

	// now, when we convert back to a, it has no decimal since a is a whole number
	b = float64(a)
	fmt.Printf("b: %8T %[1]v\n", b)
}
