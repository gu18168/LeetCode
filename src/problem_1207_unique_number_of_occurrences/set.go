package solution

func uniqueOccurrences(arr []int) bool {
	times := make(map[int]int)

	for _, num := range arr {
		times[num] += 1
	}

	set := make(map[int]struct{})

	for _, time := range times {
		if _, ok := set[time]; ok {
			return false
		}

		set[time] = struct{}{}
	}

	return true
}
