package solution

func maxOperations(nums []int, k int) int {
	occur := make(map[int]int, 0)
	result := 0

	for _, num := range nums {
		if res := k - num; occur[res] > 0 {
			result += 1
			occur[res] = occur[res] - 1
		} else {
			occur[num] = occur[num] + 1
		}
	}

	return result
}
