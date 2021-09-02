package main

import "fmt"

func main() {
	employeesByInitial := map[string]map[string]float64{
		"J": {
			"John": 2000.45,
		},
		"M": {
			"Marie": 2000.45,
		},
	}

	fmt.Println(employeesByInitial)

	for initial, employee := range employeesByInitial {
		fmt.Println(initial, employee)
	}
}
