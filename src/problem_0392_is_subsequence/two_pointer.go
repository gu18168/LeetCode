package solution

func isSubsequence(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls > lt {
		return false
	}

	is, it := 0, 0
	for ; is < ls && it < lt; it++ {
		if s[is] == t[it] {
			is++
		}
	}

	return is == ls
}
