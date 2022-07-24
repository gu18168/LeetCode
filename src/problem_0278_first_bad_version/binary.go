package solution

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	left, right := 0, n
	result := -1

	for left <= right {
		mid := (left + right) / 2

		if isBadVersion(mid) {
			result = mid

			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}
