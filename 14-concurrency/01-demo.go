package main

import "fmt"

func main() {
	panic("dummy here")
	go f1() //handing over this function execution to the scheduler
	f2()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
