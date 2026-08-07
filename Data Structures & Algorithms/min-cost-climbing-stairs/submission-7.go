func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	note := make([]int, n+1)
	note[0] = 0
	note[1] = 0

	for i := 2; i <= n; i++ {
		note[i] = min(
			note[i-1] + cost[i-1],
			note[i-2] + cost[i-2],
		)
	}

	return note[n]
}
