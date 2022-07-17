package solution

func lengthOfLongestSubstring(s string) int {
	if l := len(s); l <= 1 {
		return l
	}

	result := 0
	visited := make(map[byte]int)
	// substring: s(start, end]
	for start, end := -1, 0; end < len(s); end++ {
		ch := s[end]
		// jump to valid position
		if pos, ok := visited[ch]; ok && pos > start {
			start = pos
		}
		
		visited[ch] = end
		if l := end - start; l > result {
			result = l
		}
	}

	return result
}
