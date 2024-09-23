package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w4schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
      <img src="yyy.jpg" width="150" height="150">
  </body>
</html>`

// visits all the child elements in the html text and its siblings, this is depth first traversal
func visit(n *html.Node, pwords, ppics *int) {
	// check for tag if its an element node
	if n.Type == html.TextNode {
		(*pwords) += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		(*ppics)++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, pwords, ppics)
	}
}

// calls the visit function recusively to count all the words and pics up in the html
func count_words_and_images(doc *html.Node) (int, int) {
	var words, pics int
	visit(doc, &words, &pics)
	return words, pics
}

func main() {
	// parse into a format that the html library accepts (byte array)
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	// check for error, print it and close program
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse failed: %s\n", err)
		os.Exit(-1)
	}

	// count words and images into vars words, pics
	words, pics := count_words_and_images(doc)

	fmt.Printf("%d words and %d images\n", words, pics)
}
