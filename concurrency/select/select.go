package main

import (
	"fmt"
	"time"

	"github.com/ramonmpacheco/gohtml"
)

func theFastest(url1, url2, url3 string) string {
	c1 := gohtml.Title(url1)
	c2 := gohtml.Title(url2)
	c3 := gohtml.Title(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Too slow!!!"
	}
}

func main() {
	champ := theFastest(
		"https://www.youtube.com/",
		"https://www.google.com/",
		"https://golang.org/",
	)

	fmt.Print(champ)
}
