package solution

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

outer:
	for _, asteroid := range asteroids {
		if asteroid > 0 {
			stack = append(stack, asteroid)

			continue outer
		}

		idx := len(stack) - 1
		for ; idx >= 0 && stack[idx] > 0; idx -= 1 {
			res := stack[idx] + asteroid

			if res == 0 {
				stack = stack[:idx]

				continue outer
			} else if res > 0 {
				continue outer
			}

			stack = stack[:idx]
		}

		stack = append(stack, asteroid)
	}

	return stack
}
