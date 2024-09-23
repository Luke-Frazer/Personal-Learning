package main

import "fmt"

func main() {
	// this is null, has no length, no capacity, no values
	// if we append to this, it will allocate memory to capacity, then add at index 0
	var s []int
	// this is not null, has 0 length, 0 capacity, no values. Essentiall an array of nothing: []
	t := []int{}
	// this is not null, has 5 length, 5 capacity, and all values set to 0. Capacity is inferred by length.
	// if we append to this, it will append at index 6
	u := make([]int, 5)
	// this is not null, has 0 length, 5 capacity (essentiall 5 slots available in memory), and no values
	// if we append to this, it will append at index 0 since there is no existing length
	v := make([]int, 0, 5)

	fmt.Printf("%d, %d, %T, %t, %#[3]v\n", len(s), cap(s), s, s == nil)
	fmt.Printf("%d, %d, %T, %t, %#[3]v\n", len(t), cap(t), t, t == nil)
	fmt.Printf("%d, %d, %T, %t, %#[3]v\n", len(u), cap(u), u, u == nil)
	fmt.Printf("%d, %d, %T, %t, %#[3]v\n", len(v), cap(v), v, v == nil)
}