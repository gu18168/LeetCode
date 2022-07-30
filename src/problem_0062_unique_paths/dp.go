package solution

func uniquePaths(m int, n int) int {
	path := make([][]int, m)
	for i := range path {
		path[i] = make([]int, n)
		path[i][0] = 1
	}

	for i := range path[0] {
		path[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}

	return path[m-1][n-1]
}