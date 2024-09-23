package main

import "fmt"

func do(d func()) {
	d()
}

func main() {
	// this used to return i at the same address everytime, changed as of go 1.22
	for i := 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}

		do(v)
	}
	fmt.Println()
	
	s := make([]func(), 4)
	for j := 0; j < 4; j++ {
		s[j] = func() {
			fmt.Printf("%d @ %p\n", j, &j)
		}
	}

	// Used to return the same value at the same address due to issue in go, changed with go 1.22
	// variables in for loops instantiate every iteration
	for k := 0; k < 4; k++ {
		s[k]()
	}


}
