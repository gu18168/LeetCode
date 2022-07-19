package solution

func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		chT, okT := s2t[s[i]]
		chS, okS := t2s[t[i]]

		// okT && chT != t[i] || okS && chS != s[i]
		if okT && okS {
			if chT != t[i] || chS != s[i] {
				return false
			}
		} else if !okS && !okT {
			s2t[s[i]], t2s[t[i]] = t[i], s[i]
		} else {
			return false
		}
	}

	return true
}
