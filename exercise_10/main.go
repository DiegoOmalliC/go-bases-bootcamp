package main

import (
	"errors"
	"fmt"
)

var (
	expensive = "expensive"
	cheap     = "cheap"
	middle    = "middle"
)

type Product interface {
	calculateCost() float64
}

type CheapProduct struct {
	Price float64
}
type MiddleProduct struct {
	Price float64
}
type ExpensiveProduct struct {
	Price float64
}

func (smp *CheapProduct) calculateCost() float64 {
	return smp.Price
}
func (mdp *MiddleProduct) calculateCost() float64 {
	return mdp.Price + (mdp.Price * 0.03)
}
func (exp *ExpensiveProduct) calculateCost() float64 {
	return exp.Price + (exp.Price * 0.06) + 2500
}
func factory(category string, price float64) (Product, error) {
	switch category {
	case cheap:
		return &CheapProduct{
			Price: price,
		}, nil
	case middle:
		return &MiddleProduct{
			Price: price,
		}, nil
	case expensive:
		return &ExpensiveProduct{
			Price: price,
		}, nil
	default:
		return nil, errors.New("Not found Product")
	}
}
func main() {
	cheapProduct, err := factory("cheap", 30.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cheapProduct.calculateCost())
	}
	expensiveProduct, err := factory("expensive", 100.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(expensiveProduct.calculateCost())
	}
	middleProduct, err := factory("middle", 50.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(middleProduct.calculateCost())
	}
}
