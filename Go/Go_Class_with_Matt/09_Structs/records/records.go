package main

import "fmt"

type album3 struct {
	title string
}
type album4 struct {
	title string
}

func main() {
	// not super convenient, but can also do this method
	var album = struct {
		title  string
		artist string
		year   int
		copies int
	}{
		"My Album",
		"Me",
		2003,
		12913,
	}

	// this is why the method above is not convenient, it is an anonymous struct
	// var pAlbum *struct {
	// 	title  string
	// 	artist string
	// 	year   int
	// 	copies int
	// }

	var album2 = struct {
		title  string
		artist string
		year   int
		copies int
	}{
		"Your Album",
		"you",
		2003,
		913,
	}

	// even though these use anonymous types, they have the same struct compatibility and can be copied into each other
	album = album2

	fmt.Println(album, album2)

	var a1 = album3 {
		"Another Album",
	}
	var a2 = album4 {
		"AAAAND Another Album",
	}
	
	// a1 = a2           // Not possible in this case since these are two different named structs (album3 vs album4)
	a2 = album4(a1)   // This works however, since they have the same structure, thus they can be converted

	fmt.Println(a1, a2)


}
