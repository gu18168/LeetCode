package solution

func increasingTriplet(nums []int) bool {
	maxNums := make([]int, len(nums))

	maxVal := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		maxVal = max(maxVal, nums[i])

		maxNums[i] = maxVal
	}

	minVal := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > minVal && nums[i] < maxNums[i] {
			return true
		} else if nums[i] < minVal {
			minVal = nums[i]
		}
	}

	return false
}
