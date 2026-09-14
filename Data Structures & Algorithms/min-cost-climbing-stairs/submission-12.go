func minCostClimbingStairs(cost []int) int {
	walks := make([]int, len(cost)+1)
	walks[0], walks[1] = 0, 0

	for i := 2; i < len(walks); i++ {
		walks[i] = min(walks[i-1]+cost[i-1], walks[i-2]+cost[i-2])
	}

	return walks[len(walks)-1]
}
