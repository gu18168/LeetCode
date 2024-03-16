package solution

func mergeAlternately(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))

	minLen := min(len(word1), len(word2))

	for i := 0; i < minLen; i++ {
		result = append(result, word1[i], word2[i])
	}

	if len(word1) > minLen {
		result = append(result, word1[minLen:]...)
	} else {
		result = append(result, word2[minLen:]...)
	}

	return string(result)
}
