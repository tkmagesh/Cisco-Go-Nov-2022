package main

import "fmt"

func main() {
	//var x interface{}
	var x any
	x = 100
	x = true
	x = "This is a string"
	x = 100.99
	x = struct{}{}
	x = []int{10, 20, 30, 40}
	fmt.Println(x)

	//x = "Magesh"
	x = 100
	//y := x.(int) + 200
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// x = true
	//x = 100.999
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, 90% of x =", val*0.9)
	case struct{}:
		fmt.Println("x is an empty struct")
	default:
		fmt.Println("unknown type")
	}

}
