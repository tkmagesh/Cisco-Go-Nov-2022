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
	/*
		func(name string) {
			fmt.Printf("Hi %s, Have a nice day!\n", name)
		}("Magesh")

		func(x, y int) {
			fmt.Println(x + y)
		}(100, 200)

		result := func(x, y int) int {
			return x + y
		}(100, 200)
		fmt.Println(result)

		q, r := func(x, y int) (quotient, remainder int) {
			quotient, remainder = x/y, x%y
			return
		}(100, 7)
		fmt.Println(q, r)
	*/
}
