package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	//slice size 10, internal array size 20
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s)

	//the size will grow automatically
	s = append(s, 1)

	//copy
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)

	fmt.Println(slice2)
}
