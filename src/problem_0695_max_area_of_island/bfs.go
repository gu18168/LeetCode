package solution

func maxAreaOfIsland(grid [][]int) int {
	result := 0

	hasVisited := make([][]bool, len(grid))
	for i := range hasVisited {
		hasVisited[i] = make([]bool, len(grid[0]))
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 {
				continue
			}

			area := areaOfIsland(x, y, grid, hasVisited)

			if area > result {
				result = area
			}
		}
	}

	return result
}

func areaOfIsland(x int, y int, grid [][]int, hasVisited [][]bool) int {
	area := 0
	step := [][2]int{{x, y}}

	for len(step) != 0 {
		now := step[0]
		step = step[1:]

		x, y := now[0], now[1]

		if x > 0 && !hasVisited[y][x-1] && grid[y][x-1] == 1 {
			step = append(step, [2]int{x - 1, y})
			hasVisited[y][x-1] = true
		}
		if x < len(grid[0])-1 && !hasVisited[y][x+1] && grid[y][x+1] == 1 {
			step = append(step, [2]int{x + 1, y})
			hasVisited[y][x+1] = true
		}
		if y > 0 && !hasVisited[y-1][x] && grid[y-1][x] == 1 {
			step = append(step, [2]int{x, y - 1})
			hasVisited[y-1][x] = true
		}
		if y < len(grid)-1 && !hasVisited[y+1][x] && grid[y+1][x] == 1 {
			step = append(step, [2]int{x, y + 1})
			hasVisited[y+1][x] = true
		}

		grid[y][x] = 0
		area += 1
	}

	return area
}
