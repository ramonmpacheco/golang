package main

import (
	"fmt"
)

func main() {
	//map must be initialized
	approved := make(map[int]string)

	approved[123] = "John"
	approved[124] = "Jane"
	fmt.Println(approved)

	for key, value := range approved {
		fmt.Printf("%s (CODE: %d)\n", value, key)
	}

	fmt.Println(approved[123])
	delete(approved, 123)

	fmt.Println(approved)

	//-----------------

	employees := map[string]float64{
		"John":    2000.45,
		"Jane":    2000.45,
		"Mr Boss": 4000.454,
	}

	employees["Marie"] = 2000.45

	fmt.Println(employees)

	for name, salary := range employees {
		fmt.Printf("Name: %s, Salary: %.2f\n", name, salary)
	}

}
