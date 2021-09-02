package main

import (
	"fmt"
	"time"
)

func gradeAsLetter(g float64) string {
	var grade = int(g)

	switch grade {
	case 10:
		fallthrough // force go to next case
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid Grade!"
	}
}

func tyype(i interface{}) string {
	switch i.(type) {
	case int:
		return "Int"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Func"
	default:
		return "Unknown"
	}

}

func main() {
	fmt.Println(gradeAsLetter(9.8))
	fmt.Println(gradeAsLetter(6.8))
	fmt.Println(gradeAsLetter(2.8))

	t := time.Now()
	//look for the first case: true
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 18:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}

	fmt.Println(tyype(8))

}
