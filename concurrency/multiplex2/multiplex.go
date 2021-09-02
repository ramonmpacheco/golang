package main

import (
	"fmt"
	"time"
)

func speakUp(person string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s speaking: %d", person, i)
		}
	}()

	return c // Generator pattern
}

func putTogether(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-entry1:
				c <- s
			case s := <-entry2:
				c <- s
			}
		}
	}()

	return c
}

func main() {
	c := putTogether(
		speakUp("John"), speakUp("Mary"),
	)

	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)

}
