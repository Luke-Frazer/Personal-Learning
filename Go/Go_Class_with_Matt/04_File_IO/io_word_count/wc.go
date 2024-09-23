package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// loop through args and open the filenames
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int

		// read the bypes and print to stdout
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}

		fmt.Printf("%7s %7s %7s", "lines", "words", "chars")
		fmt.Println()
		fmt.Printf(" %7d %7d %7d %s\n", lc, wc, cc, fname)
		file.Close()
	}
}
