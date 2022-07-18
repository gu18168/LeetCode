package solution

func pivotIndex(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}

	for i := range nums[1:] {
		nums[i+1] += nums[i]
	}

	sum := nums[l-1]
	if sum-nums[0] == 0 {
		return 0
	}

	for i := 1; i < l-1; i++ {
		left, right := nums[i-1], sum-nums[i]
		if left == right {
			return i
		}
	}
	if nums[l-2] == 0 {
		return l - 1
	}

	return -1
}
