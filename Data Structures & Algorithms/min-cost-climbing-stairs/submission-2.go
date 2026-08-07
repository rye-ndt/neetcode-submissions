// reach the top, as cheap as possible 

func minCostClimbingStairs(cost []int) int {
	note := make([]int, len(cost) + 1) // off-by-one 
	// note 0 and 1 = 0 by design 

	for i := 2; i <= len(cost); i++ {
		note[i] = min(
			note[i-1] + cost[i-1],
			note[i-2] + cost[i-2],
		)
	}

	return note[len(cost)]
}
