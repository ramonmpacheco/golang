package main

import "fmt"

type item struct {
	productId int
	quantity  int
	price     float64
}

type order struct {
	userId int
	itens  []item
}

func (o order) total() float64 {
	total := 0.0
	for _, item := range o.itens {
		total += item.price * float64(item.quantity)
	}
	return total
}

func main() {
	order := order{
		userId: 1,
		itens: []item{
			{productId: 1, price: 1.99, quantity: 2},
			{productId: 2, price: 23.49, quantity: 1},
		},
	}

	fmt.Printf("Total: R$ %.2f", order.total())
}
