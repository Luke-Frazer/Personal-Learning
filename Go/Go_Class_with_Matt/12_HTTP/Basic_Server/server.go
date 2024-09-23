// Modified a bunch to work with post requests so I can mess around with http stuff outside of just a basic server
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

type todo struct {
	// just making these lowercase with the json tagging
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// tells my webpage how to render the json data
var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>`

// handles whatever route its passed into
// all output goes into the response writer, request is what we get from the client
// allows you to go to the site and it prints the url path as the "from: ", or make a post and it'll print that instead
func handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if len(body) != 0 {
		if err != nil {
			panic(err)
		}
		// include request body instead of just url path
		fmt.Fprintf(w, "Hello world from %s\n", string(body))
	} else {

		const baseUrl = "https://jsonplaceholder.typicode.com/"

		resp, err := http.Get(baseUrl + r.URL.Path[1:])

		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		defer resp.Body.Close()
		var item todo

		// even though todo struct doesn't contain userId, unmarshal will just ignore it and place the rest
		// this skips the step of needing to convert to io byte slice and then pass that in
		if err = json.NewDecoder(resp.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		// parse form text and place the response from typicode into item
		tmpl := template.New("mine")
		tmpl.Parse(form)
		tmpl.Execute(w, item)
	}
}

func main() {
	// top level route
	http.HandleFunc("/", handler)

	// opens listen socket to accept http request on port 8080, no handler
	log.Fatal(http.ListenAndServe(":8080", nil))
}
