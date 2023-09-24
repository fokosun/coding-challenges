package maths

func Factorial(x int) int {
	if x == 0 {
		return 1
	}

	if x < 0 {
		return -1
	}

	result := 1

	for i := x; i > 0; i-- {
		result *= i
	}

	return result
}
