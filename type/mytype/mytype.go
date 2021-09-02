package main

import "fmt"

type grade float64

func (g grade) between(start, end float64) bool {
	return float64(g) >= start && float64(g) <= end
}

func gradeAsLetter(g grade) string {
	if g.between(9, 10) {
		return "A"
	} else if g.between(8, 9) {
		return "B"
	} else if g.between(5, 8) {
		return "C"
	} else if g.between(3, 5) {
		return "D"
	}
	return "E"
}

func main() {
	fmt.Println(gradeAsLetter(9.8))
	fmt.Println(gradeAsLetter(6.8))
	fmt.Println(gradeAsLetter(2.8))
}
