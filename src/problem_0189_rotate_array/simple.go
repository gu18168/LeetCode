package solution

func rotate(nums []int, k int) {
	k = k % len(nums)

	for k > 0 {
		for i := 1; i < len(nums); i++ {
			nums[i], nums[0] = nums[0], nums[i]
		}

		k -= 1
	}
}
