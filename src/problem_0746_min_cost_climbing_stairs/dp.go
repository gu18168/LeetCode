package solution

func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	sum := make([]int, n)

	sum[0], sum[1] = cost[0], cost[1]
	for i := 2; i < n; i++ {
		sum[i] = min(sum[i-1], sum[i-2]) + cost[i]
	}

	return min(sum[n-1], sum[n-2])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}