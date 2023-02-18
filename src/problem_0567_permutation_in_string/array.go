package solution

func checkInclusionArray(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)

	if n1 > n2 {
		return false
	}

	chs1, chs2 := [26]int{}, [26]int{}
	for i, ch := range s1 {
		chs1[ch-'a'] += 1
		chs2[s2[i]-'a'] += 1
	}

	if chs1 == chs2 {
		return true
	}

	for i := n1; i < n2; i++ {
		chs2[s2[i]-'a'] += 1
		chs2[s2[i-n1]-'a'] -= 1

		if chs1 == chs2 {
			return true
		}
	}

	return false
}
