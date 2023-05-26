package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	tickets.ReadDataFile("tickets.csv")
	var err error
	var total int
	total, err = tickets.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Total de viajes a Brazil: ",total)
	}
	total, err= tickets.GetMornings("mañana")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Total de viajes por la mañana: ",total)
	}
	total, err = tickets.GetMornings("madrugada")
	if err != nil {
		fmt.Println(err)
    }else{
		fmt.Println("Total de viajes por la madrugada: ",total)
	}
	average, err := tickets.AverageDestination("Brazil", 10)
	if err!= nil {
		fmt.Println(err)
	}else{
		fmt.Println("Promedio de viajes por dia a Brazil:",average)
	}
}
