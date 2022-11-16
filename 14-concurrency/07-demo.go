package main

import (
	"fmt"
	"sync"
)

var no int
var mutex sync.Mutex

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
	mutex.Lock()
	no++
	mutex.Unlock()
	wg.Done()
}
