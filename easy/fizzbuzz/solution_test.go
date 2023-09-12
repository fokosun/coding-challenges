package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzzGame(t *testing.T) {
	numbers := []string{
		"1",
		"2",
		"fizz",
		"4",
		"buzz",
		"fizz",
		"7",
		"8",
		"fizz",
		"buzz",
		"11",
		"fizz",
		"13",
		"14",
		"fizz buzz",
		"16",
		"17",
		"fizz",
		"19",
		"buzz",
	}

	for i, v := range numbers {
		actual := FizzBuzz(i + 1)
		expected := v

		assert.Equal(
			t,
			expected,
			actual,
			actual+" is not the same as "+expected,
		)
	}
}
