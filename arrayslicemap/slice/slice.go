package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//not inclusive
	s2 := a2[1:3]
	fmt.Println(s2)

	// from beginning to 3 no inclusive
	//it's not a new array, but a pointer
	s3 := a2[:3]
	fmt.Println(s3)
}
