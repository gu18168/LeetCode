package solution

func moveZeroes(nums []int) {
	n := len(nums)

	zeroIdx := 0
	for idx := 0; idx < n; idx++ {
		if nums[idx] == 0 {
			continue
		}

		for ; zeroIdx < idx; zeroIdx++ {
			if nums[zeroIdx] != 0 {
				continue
			}

			nums[zeroIdx], nums[idx] = nums[idx], nums[zeroIdx]

			break
		}
	}
}
