package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[0:1] // [0, 1), slice up to len 1, but keep capacity from a, this makes b point to a's memory location
	d := a[0:1:1] // [i:j:k] len j-i, cap k-i, this specifies to slice up to len 1 and cap 1,
	c := b[0:2] // [0, 2), But this includes value 2 from a because it copies the capacity from b which copies from a

	fmt.Println("a = ", a)

	fmt.Println("b = ", b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	fmt.Println("c = ", c)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	fmt.Println("d = ", d)
	fmt.Println(len(d))
	fmt.Println(cap(d))


	fmt.Println("========================================")

	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("a[%p] = %[1]v\n", b)

	c = append(c, 5) // overrides a's 3rd value since these point to the same thing
	// however, if we set c to b[0:2:2] where the capacity is set to 2, then append would've created a new array
	fmt.Printf("a[%p] = %[1]v\n", c)
	fmt.Printf("a[%p] = %v\n", &a, a)
}
