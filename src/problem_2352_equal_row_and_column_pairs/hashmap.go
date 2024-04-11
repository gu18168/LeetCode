package solution

import (
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	n := len(grid)

	m := make(map[string]int)

	for i := 0; i < n; i++ {
		m[encodeSlice(grid[i])] += 1
	}

	result := 0

	for i := 0; i < n; i++ {
		col := make([]int, n)

		for j := 0; j < n; j++ {
			col[j] = grid[j][i]
		}

		result += m[encodeSlice(col)]
	}

	return result
}

func encodeSlice(s []int) string {
	res := strings.Builder{}

	for _, num := range s {
		res.WriteString(strconv.Itoa(num))
		res.WriteByte(',')
	}

	return res.String()
}
