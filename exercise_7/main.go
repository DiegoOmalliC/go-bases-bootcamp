package main

import (
	"errors"
	"fmt"
)

var products []Product = []Product{
	{
		ID:          1,
		Name:        "Jugo Naranja",
		Price:       20,
		Description: "N/A",
		Category:    "A",
	},
}

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p Product) save(new Product) {
	products = append(products, new)
}
func (p Product) getAll() {
	for _, v := range products {
		fmt.Println("Producto: ", v)
	}
}
func getById(ID int) (Product, error) {
	var empty Product
	for _, n := range products {
		if n.ID == ID {
			return n, nil
		}
	}
	return empty, errors.New("Not found the product")
}
func main() {
	newProduct := Product{
		ID:          2,
		Name:        "Jugo Limon",
		Price:       20,
		Description: "N/A",
		Category:    "A",
	}
	newProduct.save(newProduct)
	newProduct.getAll()
	fmt.Println(getById(newProduct.ID))
}
