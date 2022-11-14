package main

import "fmt"

func main() {
	var x int32 = 100
	var y float64
	y = float64(x) // use the type name like a function for type conversion
	fmt.Println(y)
}
