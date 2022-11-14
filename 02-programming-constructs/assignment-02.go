package main

import (
	"fmt"
	"os"
)

func main() {
	var userChoice, n1, n2, result int
	for {
		/*
			fmt.Println("1. Add")
			fmt.Println("2. Subtract")
			fmt.Println("3. Multiply")
			fmt.Println("4. Divide")
			fmt.Println("5. Exit")
			fmt.Println("Enter your choice :")
		*/
		menu := "1. Add\n2. Subtract\n3. Multiply\n4. Divide\n5. Exit\nEnter your choice : "
		fmt.Print(menu)
		fmt.Scanln(&userChoice)
		if userChoice == 5 {
			fmt.Println("Thank you!")
			os.Exit(0)
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		fmt.Println("Enter the operands : ")
		fmt.Scanln(&n1, &n2)
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Printf("Result = %d\n", result)
	}
}