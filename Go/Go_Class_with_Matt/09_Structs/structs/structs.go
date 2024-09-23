package main

import (
	"fmt"
	"time"
)

// structs are modeled after database records, allows you to store lots of data under a single record
type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {
	// map of struct points is almost always the correct way, since its hard/impossible to modify structs within a map
	employees := map[string]*Employee{}

	var e2 Employee
	// can create struct and add values separately, do not need to initialize every field
	e2.Name = "Jake"
	e2.Number = 2
	e2.Hired = time.Now().Add((-time.Hour * 24) * 60)
	employees["Jake"] = &e2

	// you can create struct types like this, but you need to specify every single variable
	employees["Luke"] = &Employee{
		"Luke",
		1,
		employees["Jake"], // cannot do this with an address (&employees["Jake"]), too dangerous since maps can change addresses often
		time.Now(),
	}

	// %+v notation is similar to %#v in that it provides more info, but not as detailed.
	fmt.Printf("All Employees:\nType: \n%T \nList:\n%+[1]v\n", employees)
	fmt.Printf("Luke: \n%+v\n", *employees["Luke"])
	fmt.Printf("Luke's Boss: \n%+v\n", *employees["Jake"])
}
