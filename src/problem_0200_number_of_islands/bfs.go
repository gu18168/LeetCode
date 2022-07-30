package solution

func numIslandsBFS(grid [][]byte) int {
	optionX := []int{0, 1, 0, -1}
	optionY := []int{1, 0, -1, 0}

	m, n := len(grid), len(grid[0])

	result := 0
	stack := make([][2]int, 0)
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == '1' {
				result++

				stack = append(stack, [2]int{x, y})
				grid[x][y] = '0'

				for len(stack) > 0 {
					now := stack[0]
					stack = stack[1:]

					for k := 0; k < 4; k++ {
						new := [2]int{now[0] + optionX[k], now[1] + optionY[k]}
						if 0 <= new[0] && new[0] < m && 0 <= new[1] && new[1] < n && grid[new[0]][new[1]] == '1' {
							stack = append(stack, new)
							grid[new[0]][new[1]] = '0'
						}
					}
				}
			}
		}
	}

	return result
}
