/*
	Write a goroutine that will generate prime numbers between the given start and end
	Receive the prime numbers and print them one at a time in the main function
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	ch := genPrimes(3, 100)
	for primeNo := range ch {
		fmt.Println(primeNo)
	}
}

func genPrimes(start, end int) <-chan int {
	ch := make(chan int)
	go func(start, end int) {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				ch <- no
			}
		}
		close(ch)
	}(start, end)
	return ch
}

func isPrime(no int) bool {
	end := int(math.Sqrt(float64(no)))
	for i := 2; i <= end; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
