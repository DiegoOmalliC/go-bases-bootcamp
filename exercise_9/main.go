package main

import "fmt"

type Alumno struct {
	Name     string
	LastName string
	Dni      string
	Date     string
}

func (a Alumno) datails() {
	fmt.Printf(" Los datos del estudiante son: \n Name: %s \n LastName: %s \n DNI: %s \n AdmissionDate: %s ", a.Name, a.LastName, a.Dni, a.Date)
}
func main() {
	alumno := Alumno{
		Name:     "Alex",
		LastName: "Jimenez",
		Dni:      "1763-34",
		Date:     "12/02/2005",
	}
	alumno.datails()
}
