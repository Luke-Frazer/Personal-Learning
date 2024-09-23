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
		
		// read the bypes and print to stdout
		data, err := io.ReadAll(file)
		
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		fmt.Println("The file has ", len(data), "bytes")
		file.Close()
	}
}
