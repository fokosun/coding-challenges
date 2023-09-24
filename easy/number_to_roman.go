package main

func NumberToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	for i := 0; num > 0; i++ {
		for num >= val[i] {
			num -= val[i]
			result += sym[i]
		}
	}

	return result
}
