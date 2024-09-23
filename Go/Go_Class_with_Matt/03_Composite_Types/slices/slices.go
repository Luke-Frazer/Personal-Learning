package main

import "fmt"

func main() {
	t := []byte("string")

	fmt.Println(len(t), t)
	fmt.Println(t[2])
	fmt.Println(t[:2]) // print up to and not including the byte at index 2
	fmt.Println(t[2:]) // print from 2 -> end
	fmt.Println(t[3:5], len(t[3:5]))
}