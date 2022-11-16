/*
	Write a goroutine that will generate prime numbers between the given start and end
	Receive the prime numbers and print them one at a time in the main function
*/

package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	ch := genPrimes(3, 100)
	for primeNo := range ch {
		fmt.Println(primeNo)
	}
}

func genPrimes(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for no := start; no <= end; no++ {
			wg.Add(1)
			go func(x int) {
				if isPrime(x) {
					ch <- x
				}
				wg.Done()
			}(no)
		}
		wg.Wait()
		close(ch)
	}()
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
