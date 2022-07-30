package solution

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	hasVisited := make([][]bool, m)
	for i := range grid {
		hasVisited[i] = make([]bool, n)
	}

	result := 0
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if !hasVisited[x][y] {
				result += island(x, y, grid, hasVisited)
			}
		}
	}

	return result
}

func island(x, y int, grid [][]byte, hasVisited [][]bool) int {
	hasVisited[x][y] = true
	if grid[x][y] == '0' {
		return 0
	}

	m, n := len(grid), len(grid[0])

	if x > 0 && !hasVisited[x-1][y] {
		island(x-1, y, grid, hasVisited)
	}
	if x < m-1 && !hasVisited[x+1][y] {
		island(x+1, y, grid, hasVisited)
	}
	if y > 0 && !hasVisited[x][y-1] {
		island(x, y-1, grid, hasVisited)
	}
	if y < n-1 && !hasVisited[x][y+1] {
		island(x, y+1, grid, hasVisited)
	}

	return 1
}
