package main

import "fmt"

func main() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}

	// this used to not work but now does as of 1.22 
	for _, item := range items {
		a = append(a, item[:])
	}

	fmt.Println(items)
	fmt.Println(a)
}