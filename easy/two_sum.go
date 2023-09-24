package main

func TwoSum(input []int, t int) []int {
	matches := make(map[int]int)
	result := make([]int, 0)

	for i := range input {
		complement := t - input[i]

		if _, ok := matches[complement]; ok {
			return append(result, matches[complement], i)
		}

		matches[input[i]] = i
	}

	return result
}
