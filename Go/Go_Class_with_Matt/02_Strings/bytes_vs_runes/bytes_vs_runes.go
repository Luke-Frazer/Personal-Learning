package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "élite"

	// this shows that the 5-character string actually has a length of 6, why?
	fmt.Printf("%8T %[1]v %d\n", s, len(s))
	fmt.Printf("%8T %[1]v\n", []rune(s))

	// we see here that this is because the string takes 6 bytes to represent itself, thus why its 6 "long"
	// there is a way to get the actual string length using the strings library
	bi := []byte(s)
	fmt.Printf("%8T %[1]v %d\n", bi, len(bi))

	// examples of how string is "described" in go
	str := "the quick brown fox"
	fmt.Println(str)

	a := len(str)                   // 19
	b := str[:3]                    // "the"
	c := str[4:9]                   // "quick"
	d := str[:4] + "slow" + str[9:] // replaces "quick"
	// str[5] = 'a'                    SYNTAX ERROR, cannot change existing string, would need to make new one + "es"
	str += "es" // now plural (copied), we can append to string, which creates a new string

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(str)

	aStr := "a string"
	x := len(aStr) // built-in = 8
	fmt.Println(x)

	fmt.Println(strings.Contains(aStr, "g")) // returns true
	fmt.Println(strings.Contains(aStr, "x")) // returns false

	fmt.Println(strings.HasPrefix(aStr, "a"))  // returns true
	fmt.Println(strings.Index(aStr, "string")) // returns 2

	aStr = strings.ToUpper(aStr) // returns "A STRING"
	fmt.Println(aStr)
}
