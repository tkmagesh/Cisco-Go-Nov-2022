package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")

	go func() {
		wg.Add(1)
		f1()
		wg.Done()
	}()

	wg.Add(1) // increment the counter (f2)
	go f2()
	fmt.Println("about to wait for the counter to become 0")
	wg.Wait() // for the counter to become 0
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f2 completed")
	wg.Done() // decrement the counter by 1
}
