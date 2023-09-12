package minmax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinPositiveNumbers(t *testing.T) {
	actual := GetMin([]int{160, 543, 23, 54, 99})
	expected := 23

	assert.Equal(t, expected, actual)
}

func TestMinZeros(t *testing.T) {
	actual := GetMin([]int{0, 0, 0})
	expected := 0

	assert.Equal(t, expected, actual)
}

func TestMinNegativeNumbers(t *testing.T) {
	actual := GetMin([]int{-23, -60, -567})
	expected := -567

	assert.Equal(t, expected, actual)
}

func TestMinAllNumbers(t *testing.T) {
	actual := GetMin([]int{-23, -60, 0, 4, 67, 567})
	expected := -60

	assert.Equal(t, expected, actual)
}

func TestMaxPositiveNumbers(t *testing.T) {
	actual := GetMax([]int{160, 543, 23, 54, 99})
	expected := 543

	assert.Equal(t, expected, actual)
}

func TestMaxZeros(t *testing.T) {
	actual := GetMax([]int{0, 0, 0})
	expected := 0

	assert.Equal(t, expected, actual)
}

func TestMaxNegativeNumbers(t *testing.T) {
	actual := GetMax([]int{-23, -60, -567})
	expected := -23

	assert.Equal(t, expected, actual)
}

func TestMaxAllNumbers(t *testing.T) {
	actual := GetMax([]int{-23, -60, 0, 4, 67, 567, -2, 0})
	expected := 567

	assert.Equal(t, expected, actual)
}
