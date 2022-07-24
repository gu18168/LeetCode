package solution

func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		num := nums[mid]

		if num > target {
			right = mid - 1
		} else if num < target {
			left = mid + 1
		} else if num == target {
			return mid
		}
	}

	return -1
}