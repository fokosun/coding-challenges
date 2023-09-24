package main

import (
	"fmt"
	"math"
)

func Sqrt(x int) float64 {
	//does not work for non-perfect square number
	if x == 0 {
		return 0
	}

	if x < 0 {
		return math.NaN()
	}

	return float64(SqrtBySubtractionMethod(x))
}

func SqrtBySubtractionMethod(primeNum int) int {
	count := 0
	i := 1
	bal := primeNum - 0
	orig := bal

	for bal > 0 {
		bal = bal - i

		i += 2
		count++

		fmt.Printf("%v %v %v %v\n", bal, orig, i, count)
	}

	return count
}
