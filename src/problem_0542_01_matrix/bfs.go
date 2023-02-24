package solution

var dirs = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func updateMatrix(mat [][]int) [][]int {
	result := make([][]int, len(mat))
	seen := make([][]bool, len(mat))

	for i := range mat {
		result[i] = make([]int, len(mat[i]))
		seen[i] = make([]bool, len(mat[i]))
	}

	queue := make([][2]int, 0)

	for x := range mat {
		for y := range mat[x] {
			if mat[x][y] == 0 {
				queue = append(queue, [2]int{x, y})
				seen[x][y] = true
			}
		}
	}

	for len(queue) != 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, dir := range dirs {
			nextX, nextY := x+dir[0], y+dir[1]

			if nextX < 0 || nextY < 0 || nextX >= len(mat) || nextY >= len(mat[0]) || seen[nextX][nextY] {
				continue
			}

			result[nextX][nextY] = 1 + result[x][y]
			queue = append(queue, [2]int{nextX, nextY})
			seen[nextX][nextY] = true
		}
	}

	return result
}
