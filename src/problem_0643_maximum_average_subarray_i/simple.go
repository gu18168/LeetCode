package solution

func findMaxAverage(nums []int, k int) float64 {
	val := 0
	for i := 0; i < k; i++ {
		val += nums[i]
	}

	maxVal := val
	for start, end := 1, k; end < len(nums); start, end = start+1, end+1 {
		val = val - nums[start-1] + nums[end]

		maxVal = max(maxVal, val)
	}

	return float64(maxVal) / float64(k)
}
