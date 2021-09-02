package main

import "math"

// Capital letters means public visibility
// lowercase or _ means private visibility at package level

// It represents a coordinate at cartesian plane
type Straight struct {
	x float64
	y float64
}

func cathetus(a, b Straight) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Calculate the linear distance between two points
func Distance(a, b Straight) float64 {
	cx, cy := cathetus(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
