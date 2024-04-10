package solution

import "sort"

func closeStrings(word1 string, word2 string) bool {
	countOfWord1 := getCount(word1)
	countOfWord2 := getCount(word2)

	for i := 0; i < 26; i++ {
		if countOfWord1[i] > 0 && countOfWord2[i] == 0 ||
			countOfWord1[i] == 0 && countOfWord2[i] > 0 {
			return false
		}
	}

	sort.Ints(countOfWord1[:])
	sort.Ints(countOfWord2[:])

	for i := 0; i < 26; i++ {
		if countOfWord1[i] != countOfWord2[i] {
			return false
		}
	}

	return true
}

func getCount(word string) [26]int {
	count := [26]int{}

	for i := range word {
		count[word[i]-'a'] += 1
	}

	return count
}
