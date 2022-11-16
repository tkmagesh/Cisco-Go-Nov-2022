package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = &sync.WaitGroup{}
	fmt.Println("main started")

	wg.Add(1)
	go f1(wg)

	wg.Add(1) // increment the counter (f2)
	go f2(wg)
	fmt.Println("about to wait for the counter to become 0")
	wg.Wait() // for the counter to become 0
	fmt.Println("main completed")
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done()
}

func f2(wg *sync.WaitGroup) {
	fmt.Println("f2 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f2 completed")
	wg.Done() // decrement the counter by 1
}
