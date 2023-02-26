package solution

func combine(n int, k int) [][]int {
	return cb(1, n+1, k)
}

func cb(start int, end int, k int) [][]int {
	result := make([][]int, 0)

	for i := start; i < end; i++ {
		if k == 1 {
			result = append(result, []int{i})

			continue
		}

		for _, v := range cb(i+1, end, k-1) {
			result = append(result, append(v, i))
		}
	}

	return result
}
