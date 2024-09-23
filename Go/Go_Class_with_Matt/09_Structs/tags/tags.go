package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	// can also do `db:<properties>` for working with databases
	Page  int      `json:"page"` // repeating the name of the var renames it in json
	Words []string `json:"words,omitempty"` // omitempty removes this in json if it doesnt get initialized
}

func main() {
	r := Response{Page: 1}
	j, _ := json.Marshal(r)

	fmt.Printf("%#v\n%v\n", r, string(j))

	var r2 Response
	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)
}
