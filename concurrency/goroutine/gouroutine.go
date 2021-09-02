package main

import (
	"fmt"
	"time"
)

func speakUp(person, text string, qtt int) {
	for i := 0; i < qtt; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i+1)
	}
}

func main() {
	// speakUp("Mary", "Why don't you talk to me?", 3)
	// speakUp("John", "I'm waiting for you to finish", 1)

	//The keyword "go" will create a go routine
	//It needs the main thread to be on
	// go speakUp("Mary", "Hi there!", 500)
	// go speakUp("John", "Hello!", 500)

	go speakUp("Mary", "Got it!", 10)
	speakUp("John", "Congrats!!", 5)

}
