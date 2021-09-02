package main

import "fmt"

func change(p1, p2 int) (second int, first int) {
	second = p2
	first = p1
	return
}

func main() {
	fmt.Print(change(7, 8))
}
