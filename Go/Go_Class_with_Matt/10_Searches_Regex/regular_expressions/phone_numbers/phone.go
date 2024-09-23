package main

import (
	"fmt"
	"regexp"
)

// () dictates capture groups, first group is 3 digits inside (), next two are 3 and 4 digits with a - between
// (123) 123-1234
const phoneRegex = `\(([[:digit:]]{3})\) ([[:digit:]]{3})-([[:digit:]]{4})`
var phoneFmt = regexp.MustCompile(phoneRegex)

func main() {
	original := "call me at (214) 514-9548 today"
	
	// Not needed, but shows what the regex matches
	match := phoneFmt.FindStringSubmatch(original)
	fmt.Printf("%q\n", match)

	// ${1} refers to the first capture group, which is defined in the phoneRegex string as the first paren group
	phoneNumber := phoneFmt.ReplaceAllString(original, "+1 ${1}-${2}-${3}")

	fmt.Println(phoneNumber)

}