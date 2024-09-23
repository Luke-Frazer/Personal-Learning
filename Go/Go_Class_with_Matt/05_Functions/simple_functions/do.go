package main

import "fmt"

func change_map_by_reference(m1 *map[int]int) {
	// formal param to a function is a local var to the func, initialized with the val passed to the actual param
	(*m1)[3] = 0
	*m1 = make(map[int]int) // remaking a map at the address of m1
	(*m1)[4] = 4
	fmt.Println("m1", m1)
}

// will not change the passed in argument m1 since it is a map
func change_map_by_value(m1 map[int]int) {
	m1[3] = 0
	m1 = make(map[int]int) 
	m1[4] = 4
	fmt.Println("m1", m1)
}

// will change the pased in argument since it is a slice
func change_slice_by_value(s1 []int) int {
	s1[1] = 0
	fmt.Println(s1)
	return s1[2]
}

func main() {
	// arrays are passed as copied variables
	m := map[int]int{4: 1, 7: 2, 8: 3}
	fmt.Println("m", m)
	change_map_by_reference(&m) // m is now changed since we passed the address in and changed the map at the address
	// all params are actually changed by copying some value, but for pointers, it points to the same thing (reference)
	fmt.Println("m", m)
	fmt.Println()

	// maps passed as values will not be changed, only changed inside of the function
	m2 := map[int]int{4: 1, 7: 2, 8: 3}
	fmt.Println("m2", m2)
	change_map_by_value(m) 
	fmt.Println("m2", m2)
	fmt.Println()

	// slices are passed by reference, this slice will get changed even though Im passing the value itself
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	v := change_slice_by_value(s)
	fmt.Println(v)
}
