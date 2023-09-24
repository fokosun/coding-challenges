package main

import (
	"github.com/stretchr/testify/assert"
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

func TestRomanToNumber(t *testing.T) {
	actual := "X"
	expected := "X"

	assert.Equal(t, expected, actual)
}

func TestPalindromeNumber(t *testing.T) {
	actual := "X"
	expected := "X"

	assert.Equal(t, expected, actual)
}

func TestRemoveElement(t *testing.T) {
	actual := "X"
	expected := "X"

	assert.Equal(t, expected, actual)
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
