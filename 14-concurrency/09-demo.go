package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var no int64

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println(no)
}

func increment(wg *sync.WaitGroup) {
	atomic.AddInt64(&no, 1)
	wg.Done()
}
