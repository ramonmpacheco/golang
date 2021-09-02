package main

import (
	"fmt"
	"strconv"
)

func main(){
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	//Warning, will print ascii table symbol 
	fmt.Println("Test " + string(123))

	// into to string
	fmt.Println("Test " + strconv.Itoa(123))

	//string to int. Atio return 2 values 1: number 2: error
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	//string to bool, can pass 0 or 1 also 
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println(b)
	}
}