package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// with * o have access directly to the value
	*n++
}

func main() {
	n := 1

	inc1(n)
	fmt.Println(n) // nothing changes

	inc2(&n) //now passing the n ref
	fmt.Println(n)
}
