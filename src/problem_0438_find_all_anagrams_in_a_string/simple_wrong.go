package solution

func findAnagramsWrong(s string, p string) []int {
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
	for i, j := 0, len(p); j <= len(s); i, j = i+1, j+1 {
		if isAnagrams(s[i:j], chs) {
			result = append(result, i)
		}
	}

	return result
}

func isAnagrams(s string, chs map[rune]int) bool {
	temp := make(map[rune]int)
	for _, ch := range s {
		expect, ok := chs[ch]
		if !ok {
			return false
		}

		if num, ok := temp[ch]; !ok {
			temp[ch] = 1
		} else if num < expect {
			temp[ch] = num + 1
		} else {
			return false
		}
	}

	return true
}
