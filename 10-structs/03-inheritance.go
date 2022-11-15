package main

import (
	"fmt"
)

type Product struct {
	Id   int
	Name string
	Cost float32
}

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	Product
	// Dummy
	Expiry string
	Name   string
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}
func main() {
	/*
		grapes := PerishableProduct{
			Product: Product{
				Id:   100,
				Name: "Grapes",
				Cost: 50,
			},
			Expiry: "2 Days",
			// Name:   "Arabian Grapes",
		}
	*/
	grapes := NewPerishableProduct(100, "Grapes", 50, "2 Days")
	fmt.Printf("%#v\n", grapes)
	fmt.Println("Accessing the fields")
	fmt.Println(grapes.Product.Id, grapes.Product.Name, grapes.Product.Cost, grapes.Expiry)
	fmt.Println(grapes.Id, grapes.Name, grapes.Cost, grapes.Expiry)
}

/*
	Write a function "Format" that will return a formatted string of the given "Product" (IMPORTANT : NOT PerishableProduct)
	Write a function ApplyDiscount that will update the cost of the given "Product" by applying the given discount (IMPORTANT : NOT PerishableProduct)

	Use the above functions with "grapes" and print the result
*/
