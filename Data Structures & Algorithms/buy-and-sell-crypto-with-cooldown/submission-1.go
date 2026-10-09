func maxProfit(prices []int) int {
	hold, sold, rest := -math.MaxInt, 0, 0

	for _, p := range prices {
		hold, sold, rest = max(hold, rest - p), hold + p, max(rest, sold)
	}

	return max(sold, rest)
}
