package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := stopSignal()
	ch := genFibonacci(stopCh)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func stopSignal() <-chan bool {
	fmt.Println("Hit ENTER to stop...")
	stopCh := make(chan bool)
	go func() {
		fmt.Scanln()
		stopCh <- true
	}()
	return stopCh
}

func genFibonacci(stopCh <-chan bool) <-chan int {
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
