package main

import "fmt"

func printResult(grade float64){
	if grade >= 7 {
		fmt.Println("Approved with grade:", grade)
	} else {
		fmt.Println("Repproved with grade:", grade)
	}
}

func gradeAsLetter(g float64) string {
	if g >= 9 && g <=10 {
		return "A"
	} else if g >= 8 && g < 9 {
		return "B"
	} else if g >= 5 && g < 8 {
		return "C"
	} else if g >= 3 && g < 5 {
		return "D"
	} 
	return "E"
}

func main() {
	printResult(7.3)
	printResult(5.3)
	
	fmt.Println(gradeAsLetter(9.8))
	fmt.Println(gradeAsLetter(6.8))
	fmt.Println(gradeAsLetter(2.8))
}