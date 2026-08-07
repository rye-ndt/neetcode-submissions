func maxProfit(prices []int) int {
	result := 0

	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			cur := prices[i] - prices[j]

			if cur > result {
				result = cur
			}
		}
	}

	if result < 0 {
		return 0
	}

	return result
}
