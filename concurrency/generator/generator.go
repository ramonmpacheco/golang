package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//Return a read-only channel
func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile(`<title>(.*?)<\/title>`)
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	t1 := title(
		"https://www.amazon.com.br/",
		"https://golang.org/",
	)

	fmt.Println("First", <-t1)
}
