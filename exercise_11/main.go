package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type MyFirstError struct {
}
type MySecondError struct {
}
type MyFourError struct {
	What int
}

func (e MyFirstError) Error() string {
	return "El salario ingresado no alcanza lo minimo posible"
}
func (e *MySecondError) Error() string {
	return "Error: el salario es menor a 10.000"
}
func (e MyFourError) Error() string {
	return fmt.Sprintf("Error: el Minimo disponible es de 150000 y el salario ingresado es: %v", e.What)
}

var MyThirdError error = errors.New("Error: el salario es menor a 10.000")

func validationError(err error) error {
	return fmt.Errorf("About Error %w", err)
}
func main() {

	var salary int = 10000
	if salary < 150000 {
		err := MyFirstError{}
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
	salary = 1500
	if salary <= 10000 {
		fmt.Println(errors.Is(validationError(&MySecondError{}), &MySecondError{}))
		fmt.Println(errors.Is(validationError(MyThirdError), MyThirdError))
	}
	if salary <= 150000 {
		err := MyFourError{
			What: salary,
		}
		fmt.Println(errors.As(err, &MyFourError{}))
		fmt.Println(err)
	}
	reader := strings.NewReader("hoal")
	b, _ := io.ReadAll(reader)
	println(string(b))

}
