package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type myError struct {
	What string
}

func (e *myError) Error() string {
	return fmt.Sprintf(e.What)
}

type Customer struct {
	Legajo  string
	Nombre  string
	Dni     string
	Phone   int
	Adrress string
}
type ManageCustomer struct {
	Customers []Customer
}

func (mc *ManageCustomer) exist(c Customer) bool {
	for _, customer := range mc.Customers {
		if customer.Dni == c.Dni {
			return true
		}
	}
	return false
}
func (mc *ManageCustomer) add(c Customer) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if mc.exist(c) == true {
		panic("Customer already exists")
	}
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if _, err := mc.validate(c); err != nil {
		panic(err.Error())
	}
	mc.Customers = append(mc.Customers, c)
	fmt.Println("Customer was added succesfully")
}
func (mc *ManageCustomer) validate(c Customer) (bool, error) {
	if c.Nombre == "" {
		return false, &myError{
			What: "there are empty values",
		}
	}
	if c.Legajo == "" {
		return false, &myError{
			What: "there are empty values",
		}
	}
	if c.Phone == 0 {
		return false, &myError{
			What: "there are empty values",
		}
	}
	if c.Adrress == "" {
		return false, &myError{
			What: "there are empty values",
		}
	}
	if c.Dni == "" {
		return false, &myError{
			What: "there are empty values",
		}
	}
	return true, nil
}

func main() {
	//OpenFile()
	//ReadFile()
	manageCustomer := &ManageCustomer{}
	customer := Customer{
		Legajo:  "Legajo",
		Nombre:  "Nombre",
		Dni:     "Dni",
		Phone:   1,
		Adrress: "Adrress",
	}
	customerTwo := Customer{
		Legajo:  "Legajo",
		Nombre:  "Nombre",
		Dni:     "Dni",
		Phone:   1,
		Adrress: "Adrress",
	}
	customerThird := Customer{
		Legajo:  "",
		Nombre:  "Nombre",
		Dni:     "Dny",
		Phone:   1,
		Adrress: "Adrress",
	}
	manageCustomer.add(customer)
	manageCustomer.add(customerTwo)
	manageCustomer.add(customerThird)

}
func OpenFile() {
	file, err := os.Open("customers.txt")
	defer func() {
		file.Close()
		fmt.Println("ejecución finalizada")
	}()
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Println(&file)
	}
}
func ReadFile() {
	file, err := os.Open("exercise_13/customers.txt")

	defer func() {
		file.Close()
		fmt.Println("ejecución finalizada")
	}()
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
