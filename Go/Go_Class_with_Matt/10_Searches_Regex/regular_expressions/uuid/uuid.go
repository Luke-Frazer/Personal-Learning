package main

import (
	"fmt"
	"regexp"
)

// ^ dictates the required start point, $ dictates the required end point, if anything is outside of that, it fails
var uuidFormat = regexp.MustCompile(`^[[:xdigit:]]{8}-[[:xdigit:]]{4}-[1-5}][[:xdigit:]]{3}-[89abAB][[:xdigit:]]{3}-[[:xdigit:]]{12}$`)

var test = []string{
	"072665ee-a034-4cc3-a2e8-9f1822c4ebbb",
	"072665ee-a034-6cc3-a2e8-9f1822c4ebbb",
	"072665ee-a034-4cc3-72e8-9f1822c4ebbb",
	"072665ee-a034-4cc3-a2e8-9f1822c4ebb",
	"072665ee-a034-4cc3-a2e8-9f1822c4ebbcb",
	"072665ee-a034-3cc3-82e8-9f1822c4ebbb",
}

func main() {
	for i, val := range test {
		if !uuidFormat.MatchString(val) {
			fmt.Printf("%d\t| %40v\t|\tfails\n", i, val)
		} else {
			fmt.Printf("%d\t| %40v\t|\tsucceeds\n", i, val)
		}
	}
}
