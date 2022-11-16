package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	no int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.no++
	}
	c.Unlock()
}

var c Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println(c.no)
}

func increment(wg *sync.WaitGroup) {
	c.Increment()
	wg.Done()
}
