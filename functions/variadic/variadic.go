package main

import "fmt"

func avg(values ...float64) float64 {
	total := 0.0

	for _, value := range values {
		total += value
	}

	return total / float64(len(values))
}

func printApproved(approved ...string) {
	fmt.Println("Approved Candidates")
	for i, app := range approved {
		fmt.Printf("%d) %s\n", i+1, app)
	}
}

func main() {
	fmt.Printf("Avg -> %.2f\n", avg(10.0, 9.0, 8.0))
	approved := []string{"John", "Jane", "Mary"}

	printApproved(approved...)
}
