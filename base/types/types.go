package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Whole numbers
	fmt.Println(1, 2, 1000)
	fmt.Println("Whole literal is", reflect.TypeOf(32000))

	// [no sign] Only positive uint8, uint16 uint32 uint64
	var b byte = 255
	fmt.Println("The byte is",  reflect.TypeOf(b))

	//[with sign] uint8, uint16 uint32 uint64
	i1 := math.MaxInt64
	fmt.Println("The int max value is", i1)
	fmt.Println("The i1 is",  reflect.TypeOf(i1))

	//Real number float32 float64
	var x float32 = 49.99
	fmt.Println("The x is",  reflect.TypeOf(x))
	fmt.Println("The literal is",  reflect.TypeOf(49.99))

	//boolean
	bo := false
	fmt.Println("The x bo",  reflect.TypeOf(bo))

	//string
	s1 := "Hello Go"
	fmt.Println("The s1 is",  reflect.TypeOf(s1)) //single quote no allowed
	fmt.Println("The string size is", len(s1))

	//string with multiple lines
	s2 := `This
	is
	a
	valid
	string`
	fmt.Println(s2)
}