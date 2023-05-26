package tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {
	ReadDataFile("../../tickets.csv")
	t.Run("should get total person for country", func(t *testing.T) {
		//Arrange
		country := "Brazil"
		expect := 45
		//Act
		totalResult, err := getTotalTickets(country)
		//Assertion
		assert.Equal(t, expect, totalResult)
		assert.Nil(t, err)
	})
	t.Run("should get zero for country that does not found", func(t *testing.T) {
		//Arrange
		country := "Guya"
		expect := 0
		//Act
		totalResult, err := getTotalTickets(country)
		//Assertion
		assert.Equal(t, expect, totalResult)
		assert.Nil(t, err)
	})
}

func TestGetMornings(t *testing.T) {
	t.Run("should get total for a time", func(t *testing.T) {
		//Arrange
		time := "mañana"
		expect := 256
		//Act
		totalResult, err := GetMornings(time)
		//Assertion
		assert.Equal(t, expect, totalResult)
		assert.Nil(t, err)
	})
	t.Run("should get total zero for a schedule that does not exist", func(t *testing.T) {
		//Arrange
		time := "medianoche"
		expect := 0
		//Act
		totalResult, err := GetMornings(time)
		//Assertion
		assert.Equal(t, expect, totalResult)
		assert.NotNil(t, err)
	})
}
func TestAverageDestination(t *testing.T) {
	t.Run("should get average destination for a time", func(t *testing.T) {
		//Arrange
		country := "Brazil"
		total := 10
		expect := 4.0
		//Act
		averageResult, err := AverageDestination(country, total)
		//Assertion
		assert.Equal(t, expect, averageResult)
		assert.Nil(t, err)
	})
	t.Run("should get average destination for a time", func(t *testing.T) {
		//Arrange
		country := "Guya"
		total := 10
		expect := 0.0
		//Act
		averageResult, err := AverageDestination(country, total)
		//Assertion
		assert.Equal(t, expect, averageResult)
		assert.Nil(t, err)
	})
	t.Run("should get error canot dive by zero ", func(t *testing.T) {
		//Arrange
		country := "Guya"
		total := 0
		expect := 0.0
		//Act
		averageResult, err := AverageDestination(country, total)
		//Assertion
		assert.Equal(t, expect, averageResult)
		assert.NotNil(t, err)
	})

}
