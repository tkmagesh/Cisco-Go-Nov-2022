package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func (emp Employee) WhoAmI() {
	fmt.Printf("I am an Employee - [%q]\n", emp.Name)
}

func (emp *Employee) AwardBonus(bonus float32) {
	emp.Salary += bonus
}

func main() {
	/*
		emp := Employee{100, "Magesh", 10000}
		emp.WhoAmI()
		emp.AwardBonus(5000)
		fmt.Println(emp.Salary) //=> 15000
	*/

	empPtr := &Employee{100, "Magesh", 10000}
	empPtr.WhoAmI()
	empPtr.AwardBonus(5000)
	fmt.Println(empPtr.Salary) //=> 15000
}
