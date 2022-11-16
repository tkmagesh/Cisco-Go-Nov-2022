package main

import (
	"fmt"
)

//share memory by communicating

//consumer
func main() {

	ch := add(100, 200)
	result := <-ch //RECEIVE operation
	fmt.Println(result)
}

//producer
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		result := x + y
		ch <- result //SEND operation
	}()
	return ch
}
