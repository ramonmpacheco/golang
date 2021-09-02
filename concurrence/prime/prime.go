package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func fprime(n int, c chan int) {
	begin := 2
	for i := 0; i < n; i++ {
		for primeCandidate := begin; ; primeCandidate++ {
			if isPrime(primeCandidate) {
				c <- primeCandidate
				begin = primeCandidate + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	go fprime(cap(c), c)

	for prime := range c {
		fmt.Printf("%d ", prime)
	}
}
