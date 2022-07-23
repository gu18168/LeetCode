package solution

func maxProfit(prices []int) int {
	result, max := 0, 0

	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		} else if profit := max - prices[i]; profit > result {
			result = profit
		}
	}

	return result
}