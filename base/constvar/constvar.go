package main

import (
	"fmt"
	"math"
)

func main() {
	// Basic const declaration sintax: const + name + type + value
	const PI float64 = 3.1415
	// Basic var declaration, the compiler will infer the type by value
	var raio = 3.2
	//Reduced var declaration 
	area := PI * math.Pow(raio, 2)
	// Declared var but not used will generate error
	fmt.Printf("The circumference area is %f \n", area)

	//Declaring const in block
	const (
		FIRST=1
		SECOND=2
	)

	//Declaring var in block
	var (
		first=1
		second=2
	)

	fmt.Println(first, second)

	//Declaring multiple var in one line
	var e, f bool = true, false
	fmt.Println(e, f)
	
	g, h, i := 2, false, "ola"
	fmt.Println(g, h, i)
}