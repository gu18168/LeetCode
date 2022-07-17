package solution

func arrayNestingInPlace(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}

	result, res := 0, length
	for i := 0; i < length; i++ {
		if nums[i] == length {
			continue
		}

		now, size := i, 0
		for nums[now] != length {
			nums[now], now = length, nums[now]
			size++
		}

		if size > result {
			result = size
		}

		res -= size
		if res < result {
			break
		}
	}

	return result
}
