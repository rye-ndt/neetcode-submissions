func maxProfit(prices []int) int {
	result := 0
	minimum := prices[0]

	for _, p := range prices {
		result = max(result, p - minimum)
		minimum = min(minimum, p)
	}

	return result
}
