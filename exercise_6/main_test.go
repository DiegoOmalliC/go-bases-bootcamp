package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxes(t *testing.T) {
	t.Run("Test taxes where salary is lower than 50000", func(t *testing.T) {
		//Arrange
		const salary float64 = 30000
		const expect float64 = 0
		//Act
		taxe := taxes(float64(salary))
		//Assertion
		assert.Equal(t, expect, taxe, "The two words should be the same.")
	})
	t.Run("Test taxes where salary sholud more than 50000 and lower than 150000", func(t *testing.T) {
		//Arrange
		const salary float64 = 60000
		const expect float64 = 10200
		//Act
		taxe := taxes(float64(salary))
		//Assertion
		assert.Equal(t, expect, taxe, "The two words should be the same.")
	})
	t.Run("Test taxes where salary is more than 150000", func(t *testing.T) {
		//Arrange
		const salary float64 = 160000
		const expect float64 = 43200
		//Act
		taxe := taxes(float64(salary))
		//Assertion
		assert.Equal(t, expect, taxe, "The two words should be the same.")
	})
}
