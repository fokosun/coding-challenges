package maths

func CalculateMean(nums []float64) float64 {
	length := float64(len(nums))
	sum := 0.0

	// Try to come up with your own solution!

	for i := range nums {
		sum += nums[i]
	}

	return sum / length
}
