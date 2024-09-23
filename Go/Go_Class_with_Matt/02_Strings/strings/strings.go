package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// need at least 3 args for this to work
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	// we want to replace an old string with a new one (EX: go run <file> old_str new_str) then type a sentance
	// can also pass in a file with `go run <file> old new < file.txt`
	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	// will continue to scan for sentances and try to replace any of the "old" word with your "new" one
	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}
}
