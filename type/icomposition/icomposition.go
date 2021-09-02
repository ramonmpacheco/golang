package main

import "fmt"

type sports interface {
	turnOnTurbo()
}

type luxurious interface {
	autoParking()
}

type sportsLuxurious interface {
	sports
	luxurious
}

type bmw7 struct{}

func (b bmw7) turnOnTurbo() {
	fmt.Println("Turbo Mode: ON")
}

func (b bmw7) autoParking() {
	fmt.Println("Starting...")
	fmt.Println("In Progress...")
	fmt.Println("Done...")
}

func main() {
	var b sportsLuxurious = bmw7{}

	b.turnOnTurbo()
	b.autoParking()
}
