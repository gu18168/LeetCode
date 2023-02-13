package solution

func isPalindromeHalf(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revX := 0

	for x > revX {
		revX = revX*10 + x%10

		x /= 10
	}

	return revX == x || revX/10 == x
}
