package solution

func gcdOfStrings(str1 string, str2 string) string {
	n, m := len(str1), len(str2)

	minLen := min(n, m)

	for i := minLen; i > 0; i-- {
		if n%i != 0 || m%i != 0 {
			continue
		}

		subStr := str1[:i]

		if gcd(str1, subStr) && gcd(str2, subStr) {
			return subStr
		}
	}

	return ""
}

func gcd(str string, subStr string) bool {
	n := len(subStr)

	for i := 0; i <= len(str)-n; i += n {
		if str[i:i+n] != subStr {
			return false
		}
	}

	return true
}
