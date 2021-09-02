package main

import "fmt"

func main(){
	x := 1
	y := 2

	//only postfix
	x++
	fmt.Println(x)
	y--
	fmt.Println(y)
	
	//Not allowed
	// fmt.Println(x++, y++)
}