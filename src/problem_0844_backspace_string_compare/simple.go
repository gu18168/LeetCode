package solution

func backspaceCompare(s string, t string) bool {
	sNew, tNew := actualString(s), actualString(t)

	return sNew == tNew
}

func actualString(s string) string {
	res := []byte{}

	end := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if end > 0 {
				end -= 1
			}
			continue
		}

		res[end] = s[i]
		end += 1
	}

	return string(res)
}