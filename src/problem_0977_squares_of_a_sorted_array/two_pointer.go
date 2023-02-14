package solution

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1

	results := make([]int, 0, len(nums))
	for left <= right {
		leftVal := nums[left] * nums[left]
		rightVal := nums[right] * nums[right]

		if leftVal >= rightVal {
			results = append(results, leftVal)
			left += 1
		} else {
			results = append(results, rightVal)
			right -= 1
		}
	}

	for i := 0; i < len(results)/2; i++ {
		results[i], results[len(results)-1-i] = results[len(results)-1-i], results[i]
	}

	return results
}
