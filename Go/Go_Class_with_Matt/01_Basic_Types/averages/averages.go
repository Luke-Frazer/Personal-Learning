package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var n int

	// Basically a while loop
	for {
		var val float64

		// fancy way of doing this, could also define _, err on the line above and if statement for err != nil below
		// need to pass address of val for Fscanln to pass the value into its location in memory
		if _, err := fmt.Fscanln(os.Stdin, &val); err != nil {  
			// IMPORTANT, ^ this is how you define a variable for only an inner scope. This is only in the if statement
			break
		}

		sum += val
		n++ // increment operator in go has to come after value and must be on its own line, not in another statement
	}

	// throw error and exit program if we have no values
	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	// after we enter values and hit ctrl + D, calculate values
	fmt.Println("The average is ", sum/float64(n))

	// To use with file: go run 'averages.go < numbers.txt' in linux
}
