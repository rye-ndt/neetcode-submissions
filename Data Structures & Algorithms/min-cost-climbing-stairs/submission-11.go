func minCostClimbingStairs(cost []int) int {
	reach := make([]int, len(cost)+1)

	for i := 2; i <= len(cost); i++ {
		reach[i] = min(reach[i-1] + cost[i-1], reach[i-2] + cost[i-2])
	}

	return reach[len(cost)]
}
