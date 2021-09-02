package main

import "fmt"

func getApprovedAmount(value int) int {
	defer fmt.Println("End!")

	if value > 5000 {
		fmt.Println("High value...")
		return 5000
	}
	fmt.Println("Low value...")
	return value
}

func main() {
	fmt.Println(getApprovedAmount(7000))
}
