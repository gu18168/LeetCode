package solution

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	prev, now := 1, 1
	for i := 2; i <= n; i++ {
		prev, now = now, prev + now
	}

	return now
}