func minCostClimbingStairs(cost []int) int {	
	walk := make([]int, len(cost)+1)
	
	for i := 2; i < len(walk); i++ {
		walk[i] = min(walk[i-1] + cost[i-1], walk[i-2] + cost[i-2])
	}

	return walk[len(walk)-1]
}