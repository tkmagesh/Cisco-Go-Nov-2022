package main

import (
	"fmt"
)

//share memory by communicating

func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch //RECEIVE operation
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result //SEND operation
}
