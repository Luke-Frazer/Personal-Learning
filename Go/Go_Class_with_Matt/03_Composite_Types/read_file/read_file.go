package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	// scans one word at a time
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		// increment the word count every time it's seen
		words[scan.Text()]++
	}

	fmt.Println(len(words), " unique words")

	// create key/value struct
	type kv struct {
		key string
		val int
	}

	// sorted key value slice
	var ss []kv

	for k, v := range words {
		// add the key/value pairs to the sorted slice
		ss = append(ss, kv{k, v})
	}

	// sort the ss slice using a passed in function as a arg that checks index pairs and compares how I specify
	sort.Slice(ss, func(i, j int) bool {
		// UPDATE: this is actuallya closure, where this function will close over ss within the Slice() method
		return ss[i].val > ss[j].val
	})

	// loop through the slice 'ss' and display the top 5 word counts
	for _, s := range ss[:5] {
		fmt.Println(s.key, "appears", s.val, "times")
	}
}
