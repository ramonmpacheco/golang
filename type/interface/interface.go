package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name     string
	lastName string
}

type product struct {
	name  string
	price float64
}

//interfaces are implemented implicitly

func (p person) toString() string {
	return p.name + " " + p.lastName
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var thing printable = person{name: "John", lastName: "Doe"}
	fmt.Println(thing.toString())
	print(thing)

	thing = product{name: "Tv", price: 1000.0} //polymorphism
	print(thing)
}
