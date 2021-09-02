package main

import "fmt"

// anonymous function
var sum = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(sum(2, 3))

	localFunc := func(a, b int) int {
		return a - b
	}

	fmt.Println(localFunc(2, 3))
}
