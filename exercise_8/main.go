package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

func NewPerson(_id int, _name string, _dateOfBirth string) *Person {
	return &Person{
		ID:          _id,
		Name:        _name,
		DateOfBirth: _dateOfBirth,
	}
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Println(e)
}
func main() {
	person_uno := NewPerson(1, "Diego", "11/11/1999")
	person_dos := NewPerson(2, "Angela", "13/11/1999")
	fmt.Println(person_uno)
	fmt.Println(person_dos)
}
