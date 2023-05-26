package main

import f "fmt"

func main(){
	suma := suma(5)
	f.Println("Suma: ",suma)
}
func suma(index int) int {
	if index == 1 {
		return index
	}
	return index + suma(index-1)
}