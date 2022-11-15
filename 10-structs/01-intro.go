package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {
	// var emp Employee
	// var emp Employee = Employee{100, "Magesh", 10000}
	// var emp Employee = Employee{Id: 100, Name: "Magesh", Salary: 10000}
	// var emp Employee = Employee{Id: 100, Salary: 10000}
	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	fmt.Println(emp)
	fmt.Println("Accessing the fields")
	fmt.Printf("Id = %d, Name = %q, Salary = %v\n", emp.Id, emp.Name, emp.Salary)

	//empPtr := &Employee{}
	//empPtr := new(Employee)
	empPtr := &Employee{
		Id:     200,
		Name:   "Suresh",
		Salary: 20000,
	}
	fmt.Println(empPtr)
	fmt.Println("Accessing the fields")
	fmt.Printf("Id = %d, Name = %q, Salary = %v\n", (*empPtr).Id, empPtr.Name, empPtr.Salary)

	/*
		e1 := Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
		e2 := Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
		fmt.Println(e1 == e2)
	*/

	e1 := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	e2 := e1
	e2.Name = "Suresh"

	fmt.Println(e1.Name, e2.Name)
}
