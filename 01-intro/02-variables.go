package main

import "fmt"

/*
var x int
var x int = 100
var x = 100
*/

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/
	/*
		var msg string = "Hello World!"
	*/
	//type inference
	/*
		var msg = "Hello World!"
	*/

	msg := "Hello World!" // := syntax can be used ONLY in a function ( NOT in the package scope)
	fmt.Println(msg)

	//x := 100 //unused variables are not allowed in a function ( but allowed at the package scope)

	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of 100 & 200 is :"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of 100 & 200 is :"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of 100 & 200 is :"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of 100 & 200 is :"
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of 100 & 200 is :"
		var result int = x + y
		fmt.Println(str, result)
	*/

	//type inference
	/*
		var x, y = 100, 200
		var str = "Sum of 100 & 200 is :"
		var result = x + y
		fmt.Println(str, result)
	*/
	/*
		var (
			x, y   = 100, 200
			str    = "Sum of 100 & 200 is :"
			result = x + y
		)
		fmt.Println(str, result)
	*/
	/*
		var x, y, str = 100, 200, "Sum of 100 & 200 is :"
		var result = x + y
		fmt.Println(str, result)
	*/

	x, y, str := 100, 200, "Sum of 100 & 200 is :"
	result := x + y
	fmt.Println(str, result)

}
