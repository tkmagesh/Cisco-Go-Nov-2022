package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genFibonacci()
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func genFibonacci() <-chan int {
	ch := make(chan int)
	stopCh := make(chan bool)
	fmt.Println("Hit ENTER to stop...")

	go func() {
		fmt.Scanln()
		stopCh <- true
	}()

	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stopCh:
				break LOOP
			default:
				ch <- x
				time.Sleep(300 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}
