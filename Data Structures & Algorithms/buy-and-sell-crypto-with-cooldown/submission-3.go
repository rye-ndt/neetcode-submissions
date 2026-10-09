func maxProfit(prices []int) int {
	own, sell, free := -math.MaxInt, 0, 0

	for _, p := range prices {
		own, sell, free = max(own, free - p), own + p, max(free, sell)
	}

	return max(sell, free)
}
