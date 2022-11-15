package main

import (
	"fmt"
	_ "modularity-demo/calculator"
	ut "modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	color.Red("app executed")
	//fmt.Println("app executed")
	/*
		fmt.Println(calculator.Divide(100, 5))
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println("Op Count = ", calculator.GetOpCount())
	*/
	fmt.Println(ut.IsEven(21))
}
