package calculator

import "fmt"

func init() {
	fmt.Println("calculator package initialized - add.go")
}

func Add(x, y int) int {
	opCount++
	return x + y
}
