package solution

func reverseString(s []byte) {
	n := len(s)

	for idx := 0; idx < n/2; idx++ {
		s[idx], s[n-1-idx] = s[n-1-idx], s[idx]
	}
}
