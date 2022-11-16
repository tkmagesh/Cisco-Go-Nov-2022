package main

import (
	"fmt"
)

func main() {
	go f1() //handing over this function execution to the scheduler
	f2()
	//time.Sleep(1 * time.Second)
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
