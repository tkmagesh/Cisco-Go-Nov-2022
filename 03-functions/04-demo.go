/* higher order functions - functions assigned to variables */
package main

import "fmt"

func main() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()

	var greet func(string)
	greet = func(name string) {
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	}
	greet("Magesh")
	func(name string) {
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	}("Magesh")

	var add func(int, int)
	add = func(x, y int) {
		fmt.Println(x + y)
	}
	add(100, 200)

	var addFn func(int, int) int
	addFn = func(x, y int) int {
		return x + y
	}
	result := addFn(100, 200)
	fmt.Println(result)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	q, r := divide(100, 7)
	fmt.Println(q, r)

}
