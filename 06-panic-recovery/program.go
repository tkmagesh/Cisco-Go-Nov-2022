package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("exiting coz of a panic", err)
			return
		}
		fmt.Println("Thank you!")
	}()
	q, r, err := divideWrapper(100, 0)
	if err != nil {
		fmt.Println("error occurred. take a different course of action:", err)
		return
	}
	fmt.Println(q, r)

	//dosomething
}

func divideWrapper(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			//convert the panic to an error
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

//3rd library
func divide(x, y int) (quotient, remainder int) {
	//programmatically raise a panic
	if y == 0 {
		panic(errors.New("divisor cannot be 0"))
	}
	quotient, remainder = x/y, x%y
	return
}
