package solution

func largestAltitude(gain []int) int {
	highest, current := 0, 0

	for _, g := range gain {
		current += g

		highest = max(highest, current)
	}

	return highest
}
