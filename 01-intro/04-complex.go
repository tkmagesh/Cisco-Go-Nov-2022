package main

import "fmt"

func main() {
	var c1 complex64
	c1 = 4 + 5i
	//fmt.Println(c1)
	fmt.Printf("%v\n", c1)
	/*
		fmt.Println("real =", real(c1))
		fmt.Println("imaginary =", imag(c1))
	*/
	/*
		fmt.Printf("real = %g\n", real(c1))
		fmt.Printf("imaginary = %g\n", imag(c1))
	*/
	//fmt.Println("real =", real(c1), "imaginary =", imag(c1))
	var complexFmt = "real = %g, imaginary = %g\n"
	fmt.Printf(complexFmt, real(c1), imag(c1))

	c2 := complex64(5 + 8i)
	c3 := c1 + c2
	fmt.Println(c3)
	var cx complex64
	fmt.Println(cx)
}
