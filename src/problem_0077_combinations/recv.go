package solution

func combineRecv(n int, k int) [][]int {
	results := make([][]int, 0)
	result := make([]int, 0)

	if n <= 0 || n < k {
		return results
	}

	var cb func(start int)
	cb = func(start int) {
		if len(result) == k {
			resultCopy := make([]int, k)
			copy(resultCopy, result)

			results = append(results, result)

			return
		}

		for i := start; i < n+1; i++ {
			if i > n-(k-len(result))+1 {
				break
			}

			result = append(result, i)
			cb(i + 1)
			result = result[:len(result)-1]
		}
	}

	cb(1)

	return results
}
