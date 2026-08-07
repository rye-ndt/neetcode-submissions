func maxProfit(prices []int) int {
	prof := 0
	minBuy := prices[0]

	for i := 0; i < len(prices); i++ {
		cur := prices[i] - minBuy 

		if cur > prof {
			prof = cur
		}

		if prices[i] < minBuy {
			minBuy = prices[i]
		}
	}

	return prof 
}
