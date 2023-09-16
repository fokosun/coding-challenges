package string_manipulation

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
		{actual: "AMAZON", expected: "NOZAMA"},
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

func TestCanReverseAGivenSliceOfStrings(t *testing.T) {
	data := []string{
		"amazon", "AMAZON", "", "12345", "M6E 0A1", "!@#$%^&*()_+", "nil",
	}

	expected := []string{
		"nil", "!@#$%^&*()_+", "M6E 0A1", "12345", "", "AMAZON", "amazon",
	}

	actual := ReverseSliceOfStrings(data)
	assert.Equal(t, expected, actual)
}

func TestStringIsAPalindrome(t *testing.T) {
	testCases := []string{
		"ada",
		"A but tuba",
		"Avid diva",
		"As I pee, sir, I see Pisa!",
		"Art, name no tub time. Emit but one mantra.",
		"Are we not pure? “No sir!” Panama’s moody Noriega brags. “It is garbage!” Irony dooms a man; a prisoner up to new era.",
		"Are we not drawn onward, we few, drawn onward to new era?",
		"Anna",
		"Anne, I vote more cars race Rome to Vienna.",
		"Are Mac ‘n’ Oliver ever evil on camera?",
		"Are we not drawn onward to new era?",
		"Air an aria.",
		"Al lets Della call Ed Stella.",
		"alula",
		"Amen icy cinema.",
		"Amore, Roma.",
		"Amy, must I jujitsu my ma?",
		"Ana",
		"Animal loots foliated detail of stool lamina.",
		"Able was I ere I saw Elba.",
		"Acrobats stab orca.",
		"Aerate pet area.",
		"Ah, Satan sees Natasha!",
		"Aibohphobia",
		"A car, a man, a maraca.",
		"A dog, a plan, a canal: pagoda.",
		"A dog! A panic in a pagoda!",
		"A lad named E. Mandala",
		"A man, a plan, a canal: Panama.",
		"A man, a plan, a cat, a ham, a yak, a yam, a hat, a canal-Panama!",
		"A new order began, a more Roman age bred Rowena.",
		"A nut for a jar of tuna.",
		"A Santa at Nasa.",
		"A Santa dog lived as a devil God at NASA.",
		"A slut nixes sex in Tulsa.",
		"A tin mug for a jar of gum, Nita.",
		"A Toyota! Race fast, safe car! A Toyota!",
		"A Toyota’s a Toyota.",
	}

	for _, word := range testCases {
		e := Palindrome(word)
		assert.True(t, e, "Asserts that the given string is a palindrome")
	}
}
