package main

import (
	"fmt"
	"time"
)

func routinne(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Done!")
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)

	go routinne(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
