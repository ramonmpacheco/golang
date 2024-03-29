package main

import "fmt"

func fact(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	result := fact(5)
	fmt.Println(result)
}
