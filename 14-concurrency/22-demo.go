package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 1)
	ch <- 100
	data := <-ch
	fmt.Println(data)
}