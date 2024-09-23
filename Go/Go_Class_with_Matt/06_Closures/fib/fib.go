package main

import "fmt"

// this is a closure, the entirety of fib() when its defined to a variable
// the closure contains the function it returns as well as the "env" such as variables and any other logic beforehand
func fib() func() int {
	a, b := 0, 1
	return func() int {
		// fibonacci sequence
		a, b = b, a+b
		return b
	}
}

func main() {
	// f saves the state of fib(), everytime it's called, the variables inside the closure change and are saved
	f := fib()
	g := fib()

	fmt.Println(g(), g(), g())

	// every loop, its just redefining x as the output of f() closure, thus it doesnt need to increment
	for x := f(); x < 100; x = f() {
		fmt.Println(x)
	}
}