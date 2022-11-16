package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	dataCtx := context.WithValue(rootCtx, "k1", "v1")
	cancelCtx, cancel := context.WithCancel(dataCtx)
	defer cancel()

	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go fn(cancelCtx, wg)
	wg.Wait()
}

func fn(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("[fn] data from context :", ctx.Value("k1"))
	no := 1
	wg.Add(1)
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	go f1(timeoutCtx, wg)
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println(no)
			no++
		}
	}
	wg.Done()
	fmt.Println("fn completed")
}

func f1(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("[f1] data from context :", ctx.Value("k1"))
	no := 10
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println(no)
			no += 10
		}
	}
	fmt.Println("f1 completed")
	wg.Done()
}
