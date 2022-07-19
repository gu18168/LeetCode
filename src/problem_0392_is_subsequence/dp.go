package solution

func isSubsequenceDP(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls > lt {
		return false
	}
	
	next := make([][26]int, lt+1)
	for i := 0; i < 26; i++ {
		next[lt][i] = lt
	}

	for i := lt - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte('a' + j) {
				next[i][j] = i
			} else {
				next[i][j] = next[i+1][j]
			}
		}
	}

	addr := 0
	for i := 0; i < ls; i++ {
		if n := next[addr][int(s[i] - 'a')]; n == lt {
			return false
		} else {
			addr = n+1
		}
	}

	return true
}
