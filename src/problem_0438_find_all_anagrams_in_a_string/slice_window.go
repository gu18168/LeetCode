package solution

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	chs := make(map[rune]int)
	for _, ch := range p {
		if num, ok := chs[ch]; ok {
			chs[ch] = num + 1
		} else {
			chs[ch] = 1
		}
	}

	result := make([]int, 0)
	for i, j := 0, 0; i <= len(s) - len(p); {
		for ; j < i + len(p); j++ {
			ch := rune(s[j])
			expect, ok := chs[ch]
			if !ok {
				break
			}

			chs[ch] = expect - 1
			// more character as expected
			if expect == 0 {
				break
			}
		}

		if j != i + len(p) {
			for ; i <= j; i++ {
				ch := rune(s[i])
				expect, ok := chs[ch]
				if ok {
					chs[ch] = expect + 1
				}

				// find new start
				if expect + 1 == 0 {
					i, j = i+1, j+1
					break
				}
			}

			if i > j {
				j = i
			}
		} else {
			result = append(result, i)
			chs[rune(s[i])] += 1
			i++
		}
	}

	return result
}
