package main

import "fmt"

func main() {
	var x int
	x = 100

	var xPtr *int
	xPtr = &x //get the address of a value
	fmt.Println(x, xPtr)

	//dereferencing
	var y int
	y = *xPtr //get the value from the address
	fmt.Println(y)

	fmt.Println(x == *(&x))

	no := 10
	fmt.Println("before incrementing, no =", no)
	fmt.Println(&no)
	increment(&no)
	fmt.Println("after incrementing, no =", no)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap( /* ? */ )
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(n *int) {
	fmt.Println(n)
	*n++
}

func swap( /* ? */ ) {

}
