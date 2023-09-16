package reverse_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestData struct {
	actual   string
	expected string
}

func TestCanReverseAGivenString(t *testing.T) {
	testCases := []TestData{
		{actual: "amazon", expected: "nozama"},
		{actual: "", expected: ""},
		{actual: "12345", expected: "54321"},
		{actual: "M6E 0A1", expected: "1A0 E6M"},
		{actual: "!@#$%^&*()_+", expected: "+_)(*&^%$#@!"},
		{actual: "nil", expected: "lin"},
	}

	for _, d := range testCases {
		actual := ReverseStr(d.actual)
		assert.Equal(t, actual, d.expected)
	}
}
