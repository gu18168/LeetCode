package solution

func runningSum(nums []int) []int {
	result := make([]int, len(nums))

	sum := 0
	for i, num := range nums {
		sum += num
		result[i] = sum
	}

	return result
}