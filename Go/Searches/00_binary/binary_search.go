package main

import (
	"fmt"
)

func binary_search(arr []int, low int, high int, val int) int {
	if high >= low {

		var mid int = (low + high) / 2

		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			return binary_search(arr, low, mid-1, val)
		} else {
			return binary_search(arr, mid+1, high, val)
		}

	}
	return -1
}

func main() {
	var arr []int = []int{1, 4, 6, 7, 19, 20, 21}
	var search int = 20
	fmt.Printf("Array: %v\nSearch value: %d\nLocation in Array: %d", arr, search, binary_search(arr, 0, len(arr)-1, search))
}
