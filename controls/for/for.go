package main

import "fmt"

func main() {
	//for while like
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//classic for
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d", i)
	}

	fmt.Println("\nMix...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

	//Infinty
	// for {

	// }
}
