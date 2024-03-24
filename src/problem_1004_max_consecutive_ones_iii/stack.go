package solution

func longestOnes(nums []int, k int) int {
	zeroStack := make([]int, 0)
	result := 0

	start := 0
	for end := 0; end < len(nums); end++ {
		if nums[end] == 1 {
			continue
		}

		zeroStack = append(zeroStack, end)

		if len(zeroStack) <= k {
			continue
		}

		result = max(result, end-start)

		start = zeroStack[0] + 1
		zeroStack = zeroStack[1:]
	}

	result = max(result, len(nums)-start)

	return result
}
