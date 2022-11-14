package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter your name:")
	//fmt.Scanln(&name)
	fmt.Scanf("%s\n", &name)
	fmt.Printf("Hi %s!\n", name)

	var n1, n2 int
	fmt.Print("Enter two numbers :")
	fmt.Scanln(&n1, &n2)
	fmt.Println(n1, n2)
}
