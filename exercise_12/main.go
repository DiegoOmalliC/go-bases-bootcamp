package main

import "fmt"

type myCustomError struct {
	What int
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("Error: El trabajador no puede trabajar menos de %v hs mensuales o el costo es negativo", e.What)
}
func main() {
	resultSalary, err := salary(400, 500)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultSalary)
	}
}
func salary(hours int, cost int) (int, error) {
	const minWorkedHours int = 80
	if (hours < minWorkedHours) || (cost < 0) {
		return 0, &myCustomError{
			What: minWorkedHours,
		}
	}
	totalSalary := float64(hours * cost)
	if totalSalary >= 150000 {
		totalSalary -= totalSalary * 0.1
	}
	return int(totalSalary), nil
}
