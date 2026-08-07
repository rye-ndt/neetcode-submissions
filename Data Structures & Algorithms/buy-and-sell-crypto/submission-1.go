func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	result := 0
	minBuy := prices[0]

	for i := 1; i < len(prices); i++ {
		cur := prices[i] - minBuy

		if cur > result {
			result = cur
		}

		if prices[i] < minBuy {
			minBuy = prices[i]
		}
	}

	if result < 0 {
		return 0
	}


	return result
}
