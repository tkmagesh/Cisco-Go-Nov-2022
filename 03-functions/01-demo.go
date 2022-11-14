package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	msg := getGreetMessage("Suresh")
	fmt.Println(msg)
	fmt.Println(add(100, 200))
	//fmt.Println(divide(100, 7))
	/*
		q, r := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
	*/

	//q,r := divide(100, 7)
	//q := divide(100, 7)
	q, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
}

//simple function
func sayHi() {
	fmt.Println("Hi there!")
}

//with argument
func greet(name string) {
	fmt.Printf("Hi %s, Have a nice day!\n", name)
}

//with argument and return result
func getGreetMessage(name string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!", name)
}

//with more than one argument
/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

//with more than one return result
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

//with name return result
/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	/*
		quotient = x / y
		remainder = x % y
	*/
	quotient, remainder = x/y, x%y
	return
}
