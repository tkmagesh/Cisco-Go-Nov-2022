/* higher order functions - returning function as result */
package main

import "fmt"

func main() {
	fn := getFn()
	fn()

	greet := getGreet()
	greet("Magesh")
}

func getFn() func() {
	return func() {
		fmt.Println("fn invoked")
	}
}

func getGreet() func(string) {
	return func(name string) {
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	}
}
