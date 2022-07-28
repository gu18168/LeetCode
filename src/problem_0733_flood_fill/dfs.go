package solution

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	hasVisited := make([][]bool, len(image))
	for i := range image {
		hasVisited[i] = make([]bool, len(image[0]))
	}

	step := make([][2]int, 1)
	step[0] = [2]int{sr, sc}

	for len(step) != 0 {
		now := step[0]
		step = step[1:]

		x, y := now[0], now[1]

		if x > 0 && !hasVisited[x-1][y] && image[x-1][y] == image[x][y] {
			step = append(step, [2]int{x - 1, y})
			hasVisited[x-1][y] = true
		}
		if x < len(image)-1 && !hasVisited[x+1][y] && image[x+1][y] == image[x][y] {
			step = append(step, [2]int{x + 1, y})
			hasVisited[x+1][y] = true
		}
		if y > 0 && !hasVisited[x][y-1] && image[x][y-1] == image[x][y] {
			step = append(step, [2]int{x, y - 1})
			hasVisited[x][y-1] = true
		}
		if y < len(image[0])-1 && !hasVisited[x][y+1] && image[x][y+1] == image[x][y]{
			step = append(step, [2]int{x, y + 1})
			hasVisited[x][y+1] = true
		}

		image[x][y] = color
	}

	return image
}
