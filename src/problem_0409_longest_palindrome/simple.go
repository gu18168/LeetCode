package solution

func longestPalindrome(s string) int {
	length := 0

	hasVisited := make(map[rune]bool, 52)
	for _, ch := range s {
		if visited, ok := hasVisited[ch]; ok && visited {
			length++
			hasVisited[ch] = false
		} else {
			hasVisited[ch] = true
		}
	}

	length *= 2
	if len(s) != length {
		length++
	}

	return length
}