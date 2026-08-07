// reach the top, as cheap as possible 

func minCostClimbingStairs(cost []int) int {
	// last move 
	// dp(i) = min of: 
	// - dp(i-1) + cost(i-1)
	// - dp(i-2) + cost(i-2)

	notebook := make([]int, len(cost)+1)

	notebook[0] = 0
	notebook[1] = 0

	for i := 2; i <= len(cost); i++ {
		notebook[i] = min(
			notebook[i-1] + cost[i-1], 
			notebook[i-2] + cost[i-2],
		)
	}

	return notebook[len(cost)]
}
