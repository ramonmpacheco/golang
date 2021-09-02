package main

import (
	"fmt"

	"github.com/ramonmpacheco/gohtml"
)

func forward(from <-chan string, to chan string) {
	for {
		to <- <-from
	}
}

func putTogether(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(entry1, c)
	go forward(entry2, c)
	return c
}

func main() {
	c := putTogether(
		gohtml.Title(
			"https://www.amazon.com.br/",
			"https://golang.org/",
		),
		gohtml.Title(
			"https://www.youtube.com/",
			"https://www.google.com/",
		),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
