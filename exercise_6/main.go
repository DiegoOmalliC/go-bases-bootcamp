package main

import (
	"errors"
	"fmt"
)

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)
const (
	typeMinimum = "minimum"
	typeAverage = "average"
	typeMaximum = "maximum"
)

type Operation func(records ...float64) float64

func main() {
	// Exercise 1
	totalTaxe := taxes(55000)
	fmt.Printf("Impuesto: $ %0.2f \n", totalTaxe)
	// Exercise 2
	resultAverage, err := average(9.0, 2.3, 6.7, 9.8)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Promedio", resultAverage)
	}
	// Exercise 3
	totalSalary, err := salary(0, "B")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Salario: $ %0.2f \n", totalSalary)
	}
	// Exercise 4
	minResultFunc, err := operation(typeMinimum)
	min := minResultFunc(2, 4, 1, 6, 4, 2, 9)
	fmt.Println("Calificacion minima: ", min)
	// Exercise 5
	var amount int
	animalDog, msg := animal("dog")
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalDog(10)
	}
	animalCat, msg := animal("cat")
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalCat(7)
	}
	animalZebra, msg := animal("zebra")
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalZebra(5)
	}
	fmt.Println("Cantidad total de alimento: ", amount)

}

/*
Exercise 1
*/
func taxes(salary float64) float64 {
	var taxe float64
	if salary > 50000 {
		taxe += salary * 0.17
		if salary > 150000 {
			taxe += salary * 0.1
		}
	}
	return taxe
}

/*
Exercise 2
*/
func average(records ...float64) (float64, error) {
	var sumationRecords float64
	for _, v := range records {
		if v < 0 {
			return 0, errors.New("Found a negative record")
		}
		sumationRecords += v
	}
	amountRecords := float64(len(records))
	return sumationRecords / amountRecords, nil
}

/*
Exercise 3
*/
func salary(workedTime int, category string) (float64, error) {
	var salary float64
	switch category {
	case "C":
		salary = (float64(workedTime) / 60) * 1000
	case "B":
		if workedTime > 0 {
			salary = ((float64(workedTime) / 60) * 1500) + (1500 * 0.20)
		} else {
			salary = 0
		}
	case "A":
		if workedTime > 0 {
			salary = ((float64(workedTime) / 60) * 1500) + (1500 * 0.50)
		} else {
			salary = 0
		}
	default:
		return 0, errors.New("Not found category")
	}
	return salary, nil
}

/*
Exercise 4
*/
func minFunc(records ...float64) float64 {
	var minRecord = records[0]
	for _, v := range records {
		if v < minRecord {
			minRecord = v
		}
	}
	return minRecord
}
func maxFunc(records ...float64) float64 {
	var maxRecord = records[0]
	for _, v := range records {
		if v > maxRecord {
			maxRecord = v
		}
	}
	return maxRecord
}
func avgRecords(records ...float64) float64 {
	var sumationRecords float64
	for _, v := range records {
		if v < 0 {
			return 0
		}
		sumationRecords += v
	}
	amountRecords := float64(len(records))
	return sumationRecords / amountRecords
}

func operation(typeOperation string) (Operation, error) {
	operations := map[string]Operation{typeMinimum: minFunc, typeMaximum: maxFunc, typeAverage: avgRecords}
	if operations[typeOperation] != nil {
		return operations[typeOperation], nil
	} else {
		return nil, errors.New("operation not found")
	}
}

/*
Exercise 5
*/
func calcDog(amount int) int {
	const foodWeight = 10000
	return amount * foodWeight
}
func calcCat(amount int) int {
	const foodWeight = 5000
	return amount * foodWeight
}
func calcHamster(amount int) int {
	const foodWeight = 250
	return amount * foodWeight
}
func calcSpider(amount int) int {
	const foodWeight = 150
	return amount * foodWeight
}

func animal(typeAnimal string) (func(int) int, error) {
	switch typeAnimal {
	case dog:
		return calcDog, nil
	case cat:
		return calcCat, nil
	case hamster:
		return calcHamster, nil
	case spider:
		return calcSpider, nil
	default:
		return nil, errors.New("Not found animal")
	}
}
