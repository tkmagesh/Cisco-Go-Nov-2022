package main

import (
	"fmt"
	"time"
)

func main() {
	timeoutCh := timeout(10 * time.Second)
	ch := genFibonacci(timeoutCh)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func timeout(delay time.Duration) <-chan time.Time {
	stopCh := make(chan time.Time)
	go func() {
		time.Sleep(delay)
		stopCh <- time.Now()
	}()
	return stopCh
}

func genFibonacci(stopCh <-chan time.Time) <-chan int {
	ch := make(chan int)
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
