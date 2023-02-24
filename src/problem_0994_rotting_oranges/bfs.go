package solution

var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := [][]int{}

	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	good := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				good++
			}
		}
	}

	if good > 0 && len(queue) == 0 {
		return -1
	}

	if len(queue) == 0 || good == 0 {
		return 0
	}

	level := 0
	for ; len(queue) != 0; level++ {
		p := [][]int{}

		for i := 0; i < len(queue); i++ {
			node := queue[i]
			x := node[0]
			y := node[1]

			if grid[x][y] == 0 || grid[x][y] == 1 || vis[x][y] {
				continue
			}

			vis[x][y] = true

			for idx := 0; idx < 4; idx++ {
				nx := x + dx[idx]
				ny := y + dy[idx]

				if nx >= 0 && ny >= 0 && nx < m && ny < n && grid[nx][ny] == 1 {
					grid[nx][ny] = 2
					p = append(p, []int{nx, ny})
				}
			}
		}

		queue = p
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return level - 1
}
