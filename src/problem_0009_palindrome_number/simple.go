package solution

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	revX, tempX := 0, x

	for tempX > 0 {
		revX = revX*10 + tempX%10

		tempX /= 10
	}

	return revX == x
}
