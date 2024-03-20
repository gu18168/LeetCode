package solution

func reverseVowels(s string) string {
	bs := []byte(s)

	for start, end := 0, len(bs)-1; end > start; start, end = start+1, end-1 {
		for ; start < end; start++ {
			if isVowel(bs[start]) {
				break
			}
		}

		if start >= end {
			break
		}

		for ; start < end; end-- {
			if isVowel(bs[end]) {
				break
			}
		}

		if start >= end {
			break
		}

		bs[start], bs[end] = bs[end], bs[start]
	}

	return string(bs)
}

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
