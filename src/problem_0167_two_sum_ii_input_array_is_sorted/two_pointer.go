package solution

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		result := numbers[left] + numbers[right]

		if result == target {
			break
		} else if result > target {
			right--
		} else {
			left++
		}
	}

	return []int{left, right}
}
