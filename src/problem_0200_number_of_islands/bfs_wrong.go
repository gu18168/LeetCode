package solution

func numIslandsBFS_wrong(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	result := 0
	stack := make([][2]int, 0)
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == '1' {
				result++

				stack = append(stack, [2]int{x, y})
				for len(stack) > 0 {
					pointX, pointY := stack[0][0], stack[0][1]
					stack = stack[1:]

					grid[pointX][pointY] = '0'

					// Since grid is set when it is processed, 
					// the same point can be added to the stack multiple times without being processed.
					// Eventually leading to OOM.
					if pointX > 0 && grid[pointX-1][pointY] == '1' {
						stack = append(stack, [2]int{pointX - 1, pointY})
					}
					if pointX < m-1 && grid[pointX+1][pointY] == '1' {
						stack = append(stack, [2]int{pointX + 1, pointY})
					}

					if pointY > 0 && grid[pointX][pointY-1] == '1' {
						stack = append(stack, [2]int{pointX, pointY - 1})
					}
					if pointY < n-1 && grid[pointX][pointY+1] == '1' {
						stack = append(stack, [2]int{pointX, pointY + 1})
					}
				}
			}
		}
	}

	return result
}
