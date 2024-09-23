package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// loop through args and open the filenames
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// assuming there are no errors, copy to the stdout for each filename
		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		file.Close()
	}
}
