package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func B() string {
	_, file, line, _ := runtime.Caller(1)
	idx := strings.LastIndexByte(file, '/')

	// retrieve the string after the / index from above and append the line number where this func is called
	return "=> " + file[idx+1:] + ":" + strconv.Itoa(line)
}

func A() string {
	// Here is where the function is called even though we are printing the A() function below
	return B()
}

func main() {
	fmt.Println(A())
}
