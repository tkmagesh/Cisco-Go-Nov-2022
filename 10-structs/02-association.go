package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
	Org    *Organization
}

type Organization struct {
	Name string
	City string
}

func main() {
	cisco := Organization{
		Name: "Cisco",
		City: "Bangalore",
	}

	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
		Org:    &cisco,
	}

	fmt.Println(emp.Org.Name, emp.Org.City) // => Cisco, Bangalore
	cisco.City = "Chennai"
	fmt.Println(emp.Org.Name, emp.Org.City) // =>
}
