package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Sub =>", a-b)
	fmt.Println("Div =>", a/b)
	fmt.Println("Mult =>", a*b)
	fmt.Println("Mod =>", a%b)
	
	//bitwise
	fmt.Println("E =>", a&b) // binary value 3=11 & 2=10 = 10
	fmt.Println("OR =>", a|b)
	fmt.Println("XOR =>", a^b)

	c := 3.0
	d := 2.0

	fmt.Println("Greater value =>", math.Max(float64(a), float64(b)))
	fmt.Println("Smaller value =>", math.Min(float64(c), float64(d)))
	fmt.Println("Expo =>", math.Pow(c, d))
	

}