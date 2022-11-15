package main

import (
	"fmt"
)

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v", product.Id, product.Name, product.Cost)
}

func (product *Product) ApplyDiscount(discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
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

	grapes := NewPerishableProduct(100, "Grapes", 50, "2 Days")

	//fmt.Println(Format(grapes.Product))
	//fmt.Println(grapes.Product.Format())
	fmt.Println(grapes.Format())
	fmt.Println("After applying 10% discount")
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())

	/*
		fmt.Println("After applying 10% discount")
		ApplyDiscount(&grapes.Product, 10)
		fmt.Println(FormatPerishableProduct(grapes))
	*/
}

/*
	Write a function "Format" that will return a formatted string of the given "Product" (IMPORTANT : NOT PerishableProduct)
	Write a function ApplyDiscount that will update the cost of the given "Product" by applying the given discount (IMPORTANT : NOT PerishableProduct)

	Use the above functions with "grapes" and print the result

*/

/* func FormatPerishableProduct(pp PerishableProduct) string {
	return fmt.Sprintf("%s, Expiry = %q", Format(pp.Product), pp.Expiry)
} */
