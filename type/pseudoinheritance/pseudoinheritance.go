package main

import "fmt"

type car struct {
	name         string
	maxSpeed     int
	currentSpeed int
}

type ferrari struct {
	car
	turboOn bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.currentSpeed = 0
	f.turboOn = true

	fmt.Printf(
		"Does the ferrari %s have the turbo on? %v\n",
		f.name, f.turboOn,
	)
}
