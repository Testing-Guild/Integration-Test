package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Calculate(x int, y int) int {
	return x + y
}
func TestCalculate(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{"Basic Addition", 2, 2, 4},
		{"Negative Numbers", -1, 1, 0},
		{"Large Numbers", 10000, 5000, 15000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Calculate(test.x, test.y) 
			assert.Equal(t, test.expected, actual, "Expected %d + %d to equal %d", test.x, test.y, test.expected)
		})
	}
	assert.NotEqual(t, Calculate(2, 2), 5, "2 + 2 should not equal 5")
}
