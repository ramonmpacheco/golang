package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s).Intn(10)
}

func main() {
	//init a scoped var i
	if i := randomNumber(); i > 5 {
		fmt.Println("Win!!")
	} else {
		fmt.Println("Lose!!")
	}

	//i is not accessible
	// fmt.Println(i)
}
