package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(n int) string {
	//fmt.Println(n)
	if n%15 == 0 {
		return "fizz buzz"
	} else if n%5 == 0 {
		return "buzz"
	} else if n%3 == 0 {
		return "fizz"
	} else {
		return strconv.Itoa(n)
	}
}
