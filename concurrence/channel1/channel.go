package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 //sending data to channel (write)
	<-ch    //receiving data from channel (read)

	ch <- 2
	fmt.Println(<-ch)
}
