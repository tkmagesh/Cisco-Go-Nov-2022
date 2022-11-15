package calculator

import (
	"fmt"
	_ "modularity-demo/calculator/utils"
)

var opCount int

func init() {
	fmt.Println("calculator package initialized - 1")
}

func init() {
	fmt.Println("calculator package initialized - 2")
}

func Divide(x, y int) int {
	opCount++
	return x / y
}

func GetOpCount() int {
	return opCount
}
