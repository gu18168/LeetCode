package solution

func productExceptSelf(nums []int) []int {
	results := make([]int, len(nums))
	for i := range results {
		results[i] = 1
	}

	prev := 1
	for i := 1; i < len(nums); i++ {
		prev *= nums[i-1]

		results[i] = results[i] * prev
	}

	prev = 1
	for i := len(nums) - 2; i >= 0; i-- {
		prev *= nums[i+1]

		results[i] = results[i] * prev
	}

	return results
}
