package maths

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindRootsOfAGeometricalEquation(t *testing.T) {
	x1, x2 := FindRoots(2, 10, 8)

	expected1 := -1.0
	expected2 := -4.0

	assert.Equal(t, x1, expected1)
	assert.Equal(t, x2, expected2)
}

func TestCalculateMean(t *testing.T) {
	actual := fmt.Sprintf("%.2f", CalculateMean([]float64{1.34, 2.345}))
	expected := fmt.Sprintf("%.2f", 1.8425)

	assert.Equal(t, expected, actual)
}
