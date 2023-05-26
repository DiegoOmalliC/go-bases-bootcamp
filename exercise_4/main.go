package main

import f "fmt"

func main() {
	SpellWord()
	Bank()
	Month(3)
	Ages()
}
func SpellWord() {
	word := "aplicaciones"
	f.Println("Size: ", len(word))
	f.Println("Letters:")
	for _, v := range word {
		f.Printf("%v, ", string(v))
	}
	f.Println()
}
func Bank() {
	const topSalary float64 = 100000
	type Employee struct {
		Age       int
		Active    bool
		Antiquity int
		Salary    float64
	}

	daniel := Employee{
		Age:       25,
		Active:    true,
		Antiquity: 3,
		Salary:    230000.00,
	}
	if daniel.Age < 22 || !daniel.Active || daniel.Antiquity <= 1 {
		f.Println("Your not qualifies to receive the loan")
		return
	}
	f.Println("Your qualifies to receive the loan")
	if daniel.Salary > topSalary {
		f.Println("We wont charge interest")
	}

}
func Month(key int) {
	var months = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}
	f.Println(months[key])
}
func Ages() {
	var employees = map[string]int{
		"Benjamin": 28,
		"Nahuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}
	olderEmployees := 0
	for _, employee := range employees {
		if employee <= 21 {
			continue
		}
		olderEmployees++
	}
	f.Println("La edad de Benjamin es: ", employees["Benjamin"])
	f.Println("Empleados mayores de 21 años: ", olderEmployees)
	employees["Federico"] = 32
	delete(employees, "Pedro")
	f.Println(employees)
}
