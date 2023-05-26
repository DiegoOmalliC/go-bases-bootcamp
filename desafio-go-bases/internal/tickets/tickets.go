package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id      int
	Name    string
	Email   string
	Country string
	Date    int
	Price   int
}

var (
	earlyMorning = "madrugada"
	morning      = "mañana"
	afternoon    = "tarde"
	night        = "noche"
)
var Tickets = []Ticket{}

func ReadDataFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
		file.Close()
	}()
	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		var ticket Ticket
		if value, err := strconv.Atoi(line[0]); err == nil {
			ticket.Id = value
		}
		ticket.Name = line[1]
		ticket.Email = line[2]
		ticket.Country = line[3]
		if value, err := strconv.Atoi(strings.Split(line[4], ":")[0]); err == nil {
			ticket.Date = value
		}
		if value, err := strconv.Atoi(line[5]); err == nil {
			ticket.Price = value
		}
		Tickets = append(Tickets, ticket)
	}
}

func getTotalByDate(minDate, maxDate int) int {
	total := 0
	for _, ticket := range Tickets {
		if ticket.Date >= minDate && ticket.Date <= maxDate {
			total++
		}
	}
	return total
}

// ejemplo 1
func getTotalTickets(destination string) (int, error) {
	if len(Tickets) <= 0 {
		return 0, errors.New("Not found data")
	}
	total := 0
	for _, v := range Tickets {
		if v.Country == destination {
			total++
		}
	}
	return total, nil
}

// ejemplo 2
func GetMornings(time string) (int, error) {
	if len(Tickets) <= 0 {
		return 0, errors.New("Not found data")
	}
	switch time {
	case earlyMorning:
		return getTotalByDate(0, 6), nil
	case morning:
		return getTotalByDate(7, 12), nil
	case afternoon:
		return getTotalByDate(13, 19), nil
	case night:
		return getTotalByDate(20, 23), nil
	default:
		return 0, errors.New("Not found time")
	}
}

// ejemplo 3
func AverageDestination(destination string, total int) (float64, error) {
	totalCountryTickets, err := getTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	if total <= 0 || totalCountryTickets <= 0 {
		return 0, errors.New("Cannot dive by zero or negative numbers")
	}
	return float64(totalCountryTickets / total), nil
}
