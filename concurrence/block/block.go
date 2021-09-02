package main

import (
	"fmt"
	"time"
)

func routinne(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // block action
	fmt.Println("After channel data being read")
}

func main() {
	c := make(chan int) // without buffer

	go routinne(c)

	fmt.Println(<-c) //block action
	fmt.Println("Being read")
	fmt.Println(<-c)   // deadlock
	fmt.Println("End") // never reached
}
