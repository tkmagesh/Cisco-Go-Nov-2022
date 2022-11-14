package main

import "fmt"

func main() {
	var c1 complex64
	c1 = 4 + 5i
	fmt.Println(c1)
	fmt.Println("real =", real(c1))
	fmt.Println("imaginary =", imag(c1))
	c2 := complex64(5 + 8i)
	c3 := c1 + c2
	fmt.Println(c3)
}
