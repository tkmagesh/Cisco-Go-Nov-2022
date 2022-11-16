package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(7 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(5 * time.Second)
		data3 := <-ch3
		fmt.Println(data3)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 200
	}()

	/*
		The select statement lets a goroutine wait on multiple communication operations.

		A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
	*/
	for i := 0; i < 3; i++ {
		select {
		case ch3 <- 300:
			fmt.Println("Data sent to ch3")
		case data1 := <-ch1:
			fmt.Println(data1)
		case data2 := <-ch2:
			fmt.Println(data2)
			/* default:
			fmt.Println("no channel operation succeeded") */
		}
	}
}
