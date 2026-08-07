func maxProfit(prices []int) int {
	result := 0
	buy := prices[0]

	for i := 0; i < len(prices); i++ {
		cur := prices[i]

		if cur - buy > result {
			result = cur - buy
		}

		if cur < buy {
			buy = cur
		}
	}

	return result
}
