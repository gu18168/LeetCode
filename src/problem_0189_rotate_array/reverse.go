package solution

func rotateReverse(nums []int, k int) {
	k %= len(nums)

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for left, n := 0, len(nums); left < n/2; left++ {
		nums[left], nums[n-1-left] = nums[n-1-left], nums[left]
	}
}
