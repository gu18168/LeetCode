package solution

func removeStars(s string) string {
	bStack := make([]byte, 0)

	for i := range s {
		c := s[i]

		if c != '*' {
			bStack = append(bStack, c)
		} else if len(bStack) != 0 {
			bStack = bStack[:len(bStack)-1]
		}
	}

	return string(bStack)
}
