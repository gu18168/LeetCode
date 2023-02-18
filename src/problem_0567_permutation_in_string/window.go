package solution

func checkInclusion(s1 string, s2 string) bool {
	chs := make([]int, 26)
	for _, ch := range s1 {
		chs[ch-'a'] += 1
	}

	left := 0
	for right := 0; right < len(s2); right++ {
		idx := s2[right] - 'a'

		if chs[idx] -= 1; chs[idx] == -1 {
			for ; left <= right; left++ {
				if chs[idx] == 0 {
					break
				}

				chs[s2[left]-'a'] += 1
			}
		}

		if right-left+1 == len(s1) {
			return true
		}
	}

	return false
}
