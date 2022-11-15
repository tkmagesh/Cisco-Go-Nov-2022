/*
Refactor the solution for Assignment-01
	write a function genPrimes that generates the prime numbers from the given start to end and returns the primeNos
	the main function receives the prime nos and prints them
*/

package main

import "fmt"

func main() {
	primes := genPrimes(3, 100)
	fmt.Println(primes)

}

func genPrimes(start, end int) []int {
	primes := make([]int, 0)
LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		primes = append(primes, no)
	}
	return primes
}
