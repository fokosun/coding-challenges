package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type TestData struct {
	actual   []int
	expected []int
	descr    string
}

func TestHelloWorld(t *testing.T) {
	actual := Greeting()
	expected := "Hello world!"

	assert.Equal(t, expected, actual)
}

func TestNumberToRoman(t *testing.T) {
	actual := NumberToRoman(1994)
	expected := "MCMXCIV"

	assert.Equal(t, expected, actual)
}

func TestRomanToNumber(t *testing.T) {
	actual := RomanToNumber("MCMXCIV")
	expected := 1994

	assert.Equal(t, expected, actual)
}

func TestPalindromeNumber(t *testing.T) {
	testTrueData := []int{
		121, 1, 12121,
	}

	testFalseData := []int{
		-121, 10, 5001,
	}

	for _, i := range testTrueData {
		assert.True(t, PalindromeNumber(i))
	}

	for _, i := range testFalseData {
		assert.False(t, PalindromeNumber(i))
	}
}

func TestRemoveElement(t *testing.T) {
	actual1 := RemoveElement([]int{3, 2, 2, 3}, 3)
	expected1 := 2

	assert.Equal(t, actual1, expected1)

	actual2 := RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	expected2 := 5

	assert.Equal(t, expected2, actual2)
}

func TestTwoSum(t *testing.T) {
	testCases := []TestData{
		{actual: TwoSum([]int{2, 7, 11, 15}, 9), expected: []int{0, 1}, descr: "case sum found"},
		{actual: TwoSum([]int{-2, -7, -11, 15}, 4), expected: []int{2, 3}, descr: "case sum found"},
		{actual: TwoSum([]int{0, 0}, 0), expected: []int{0, 1}, descr: "case sum found"},
		{actual: TwoSum([]int{0, 0, 0, 0, 0}, 0), expected: []int{0, 1}, descr: "case sum found"},
		{actual: TwoSum([]int{0}, 0), expected: []int{}, descr: "case sum not found"},
		{actual: TwoSum([]int{2, 7, 11, 15}, 90), expected: []int{}, descr: "case sum not found"},
	}

	for _, d := range testCases {
		assert.Equal(t, d.actual, d.expected)
	}
}

func TestSqrtCasePerfectSquares(t *testing.T) {
	sqrs := make(map[int]int)

	for i := 1; i <= 100; i++ {
		sqrs[i*i] = i
	}

	for sq, sqrt := range sqrs {
		actual := Sqrt(sq)
		expected := float64(sqrt)

		assert.Equal(t, expected, actual)
	}
}

func TestSqrtCaseZero(t *testing.T) {
	actual := Sqrt(0)
	expected := float64(0)

	assert.Equal(t, expected, actual)
}

func TestSqrtCaseImaginaryNumbers(t *testing.T) {
	imgnry := make(map[int]int, 100)

	for i := 1; i <= 100; i++ {
		imgnry[i] = -i
	}

	fmt.Println(imgnry)

	for _, c := range imgnry {
		assert.True(t, math.IsNaN(Sqrt(c)))
	}
}

func TestSqrtCaseNonPerfectSquares(t *testing.T) {
	nums := map[int]float64{
		//2: 1.4142,
		//3: 1.7221,
		//5: 2.2361,
		//7: 2.6458,
		17: 4.1231,
	}

	for n, sqrt := range nums {
		actual := Sqrt(n)
		expected := sqrt

		assert.Equal(t, expected, actual)
	}
}
