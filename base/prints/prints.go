package main

import "fmt"

func main() {
	fmt.Print("Same")
	fmt.Print(" line")

	fmt.Println("")
	fmt.Println("New line")

	x := 3.141516

	fmt.Println("The x value is", x)
	fmt.Printf("The x value is %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "Text"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d,)
}