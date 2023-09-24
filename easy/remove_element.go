package main

func RemoveElement(nums []int, val int) int {
	newLength := 0

	// Iterate through the array using two pointers technique.
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			// If the current element is not equal to the target value, copy it to the new position.
			nums[newLength] = nums[i]
			newLength++
		}
	}

	return newLength
}
