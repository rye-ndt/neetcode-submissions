func minCostClimbingStairs(cost []int) int {
	note := make([]int, len(cost)+1)
	note[0] = 0
	note[1] = 0 // takes nothing to climb to here 

	for i := 2; i <= len(cost); i++ {
		note[i] = min(
			note[i-1] + cost[i-1], 
			note[i-2] + cost[i-2],
		)
	}

	return note[len(cost)]
}
