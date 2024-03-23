package solution

func maxVowels(s string, k int) int {
	num := 0
	for i := range s[:k] {
		num += isVowel(s[i])
	}

	result := num
	for end := k; end < len(s); end++ {
		num = num - isVowel(s[end-k]) + isVowel(s[end])

		result = max(result, num)
	}

	return result
}

func isVowel(b byte) int {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return 1
	default:
		return 0
	}
}
