package deduplicate_slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeDuplicateSliceCapitalized(t *testing.T) {
	actual := UniqueNames([]string{"Ava", "Emma", "Olivia"}, []string{"Olivia", "Sophia", "Emma", "George"})
	expected := []string{"Ava", "Emma", "Olivia", "Sophia", "George"}

	for _, v := range actual {
		assert.Contains(t, expected, v)
	}
}

func TestDeDuplicateSliceAllLowerCases(t *testing.T) {
	actual := UniqueNames([]string{"ava", "emma", "olivia"}, []string{"olivia", "sophia", "emma", "george"})
	expected := []string{"ava", "emma", "olivia", "sophia", "george"}

	for _, v := range actual {
		assert.Contains(t, expected, v)
	}
}

func TestDeDuplicateSliceAllUpperCases(t *testing.T) {
	actual := UniqueNames([]string{"AVA", "EMMA", "OLIVIA"}, []string{"OLIVIA", "SOPHIA", "EMMA", "GEORGE"})
	expected := []string{"AVA", "EMMA", "OLIVIA", "SOPHIA", "GEORGE"}

	for _, v := range actual {
		assert.Contains(t, expected, v)
	}
}

func TestDeDuplicateSliceCombinedCases(t *testing.T) {
	actual := UniqueNames([]string{"AVA", "eMMA", "olivia"}, []string{"Olivia", "SopHIa", "EMMa", "GEoRgE"})
	expected := []string{"AVA", "EMMa", "Olivia", "SopHIa", "GEoRgE"}

	for _, v := range actual {
		assert.Contains(t, expected, v)
	}
}
