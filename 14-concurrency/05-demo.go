package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hit ENTER to start")
	fmt.Scanln()
	var wg = &sync.WaitGroup{}
	fmt.Println("main started")
	noOfGoroutines := flag.Int("count", 0, "Number of goroutines to spawn")
	flag.Parse()
	rand.Seed(7)
	for i := 1; i <= *noOfGoroutines; i++ {
		delay := time.Duration(rand.Intn(10)) * time.Second
		wg.Add(1)
		go fn(wg, i, delay)
	}
	wg.Wait() // for the counter to become 0
	fmt.Println("main completed")
	fmt.Println("Hit ENTER to shutdown")
	fmt.Scanln()
}

func fn(wg *sync.WaitGroup, id int, delay time.Duration) {
	fmt.Printf("f1[%d] started\n", id)
	time.Sleep(delay)
	fmt.Printf("f1[%d] completed\n", id)
	wg.Done()
}
