package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	// type inference
	i := 3

	// add assignment
	i += 3
	// sub assignment
	i -= 3
	// div assignment
	i /= 2
	// mult assignment
	i *= 2
	// mod assignment
	i %= 2
	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}