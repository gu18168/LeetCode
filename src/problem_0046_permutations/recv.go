package solution

func permute(nums []int) [][]int {
	results := [][]int{}
	result := []int{}

	visited := map[int]bool{}

	var recv func()
	recv = func() {
		if len(result) == len(nums) {
			resultCopy := make([]int, len(result))
			copy(resultCopy, result)

			results = append(results, resultCopy)

			return
		}

		for _, num := range nums {
			if visited[num] {
				continue
			}

			result = append(result, num)
			visited[num] = true

			recv()

			result = result[:len(result)-1]
			visited[num] = false
		}
	}

	recv()

	return results
}
