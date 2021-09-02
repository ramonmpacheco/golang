package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

// function with receiver
func (p product) priceWithDiscount() float64 {
	return p.price - p.price*p.discount
}

func main() {
	var product1 = product{
		name:     "Pencil",
		price:    1.79,
		discount: 0.05,
	}

	fmt.Printf("%.2f", product1.priceWithDiscount())
}
