package minmax

func GetMin(nums []int) int {
	values := nums

	minimum := values[0]

	// Try to come up with your own solution

	//for _, value := range values[1:] {
	//	if value <= minimum {
	//		minimum = value
	//	}
	//}

	return minimum
}

func GetMax(nums []int) int {
	values := nums

	maximum := values[0]

	// Try to come up with your own solution

	//for _, value := range values[1:] {
	//	if value > maximum {
	//		maximum = value
	//	}
	//}

	return maximum
}
