package solution

func fib(n int) int {
	if n < 2 {
		return n
	}

	prev, now := 0, 1
	for i := 2; i <= n; i++ {
		prev, now = now, prev + now
	}

	return now
}