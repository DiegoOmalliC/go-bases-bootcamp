package main

import f "fmt"

func main(){

	arreglo := [3]int{1,2,3}
	slice := arreglo[0:len(arreglo)]
	f.Printf("arreglo: %T , slice: %T \n",arreglo,slice)
	f.Println(slice)
	for _, v := range slice {
		switch v {
		case 1:
			f.Println(v)
		case 2:
			f.Println(v)
		default:
			f.Println("Out cases")
		}

	}
	chain := "hola"
	for i,v := range chain {
		f.Println(i,string(v))
	}
	f.Println(suma(5))
}
func suma(index int) int {
	if index == 1 {
		return index
	}
	return index + suma(index-1)
}