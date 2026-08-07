// reach the top, as cheap as possible 

func minCostClimbingStairs(cost []int) int {
	prev1 := 0
	prev2 := 0

	for i := 2; i <= len(cost); i++ {
		cur := min(
			prev1 + cost[i-1],
			prev2 + cost[i-2],
		)

		prev2 = prev1
		prev1 = cur
	}

	return prev1
}
