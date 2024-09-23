package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// structure of post request
type Post struct {
	// just making these lowercase with the json tagging
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func post_string(pString string) {

	// my endpoint, just using localhost for this
	postUrl := "http://localhost:8080/"

	// make the request
	r, err := http.NewRequest("POST", postUrl, strings.NewReader(pString))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "text/plain")

	// get the address of a new client and make the request using client.do
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(r)

	if err != nil {
		if strings.Contains(err.Error(), "connectex: No connection could be made") {
			fmt.Println("The connection was not made, is the server running?")
			os.Exit(-1)
		} else {
			panic(err)
		}
	}

	// this will close the client at the end of the function
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		responseString := string(body)

		fmt.Println("POST REQUEST:")
		fmt.Println("Status: ", res.Status)
		fmt.Println("Response Body: ", responseString)
	}
}

func main() {
	if strings.Contains(strings.ToLower(os.Args[1]), "post") {
		post_string(os.Args[2])
	}
}
