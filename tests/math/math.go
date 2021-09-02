package math

import (
	"fmt"
	"strconv"
)

//Calculate avarage
func Avg(numbers ...float64) float64 {
	result := 0.0

	for _, i := range numbers {
		result += i
	}

	avg := result / float64(len(numbers))
	avgRound, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 64)

	return avgRound
}