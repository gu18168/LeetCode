package solution

func longestSubarray(nums []int) int {
	result := 0

	left, lSum, rSum := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		rSum += 1 - nums[right]

		for rSum-lSum > 1 {
			lSum += 1 - nums[left]
			left += 1
		}

		result = max(result, right+1-left-rSum+lSum)
	}

	result = max(result, len(nums)-left-rSum+lSum)

	return result - 1 + rSum - lSum
}
