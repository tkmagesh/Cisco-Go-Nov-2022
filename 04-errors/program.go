package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be 0")

func main() {

	quotient, remainder, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Do not attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
	fmt.Printf("Quotient = %d, Remainder = %d\n", quotient, remainder)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := fmt.Errorf("divisor error, divisor = %d", y)
		return 0, 0, err
	}
	return x / y, x % y, nil
}
*/

/*
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		//err = fmt.Errorf("divisor error, divisor = %d", y)
		err = errors.New("divisor cannot be 0")
		return
	}
	quotient, remainder = x/y, x%y
	return
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		//err = fmt.Errorf("divisor error, divisor = %d", y)
		err = DivideByZeroError
		return
	}
	quotient, remainder = x/y, x%y
	return
}
