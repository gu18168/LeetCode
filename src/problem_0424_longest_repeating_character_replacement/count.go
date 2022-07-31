package solution

func characterReplacement(s string, k int) int {
	result := 0

	count := [26]int{}
	max := 0
	for i, j := 0, 0; j < len(s); j++ {
		ch := s[j] - 'A'
		count[ch] += 1

		if count[ch] > max {
			max = count[ch]
		}

		for ;j - i + 1 > k + max; i++ {
			ch = s[i] - 'A'
			count[ch] -= 1

			max = 0
			for k := 0; k < 26; k++ {
				if count[k] > max {
					max = count[k]
				}
			}
		}

		if j - i + 1 > result {
			result = j - i + 1
		}
	}

	return result
}