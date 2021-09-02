package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastName string
}

func (p person) getFullName() string {
	return p.name + " " + p.lastName
}

func (p *person) setFullName(fullName string) {
	parts := strings.Split(fullName, " ")

	p.name = parts[0]
	p.lastName = parts[1]
}

func main() {
	p1 := person{name: "John", lastName: "Doe"}

	fmt.Println(p1.getFullName())

	p1.setFullName("Jane Doe")

	fmt.Println(p1.getFullName())
}
