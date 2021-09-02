package main

import "fmt"

// args: name + type
// return int
func sum(a int, b int) int {
	return a + b
}

//return  void
func print(value int) {
	fmt.Println(value)
}

func main() {
	println(sum(4, 4))
}