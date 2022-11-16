package main

import (
	"fmt"
	"sync"
)

//share memory by communicating

func main() {
	wg := &sync.WaitGroup{}
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)
	wg.Add(1)
	go add(100, 200, wg, ch)
	wg.Wait()
	result := <-ch //RECEIVE operation
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result //SEND operation
	wg.Done()
}
