func climbStairs(n int) int {
	if n < 3 { return n }

	ways := make([]int, n)
	ways[0] = 1 
	ways[1] = 2 

	for i := 2; i < len(ways); i++ {
		ways[i] = ways[i-2] + ways[i-1]
	}

	return ways[len(ways)-1]
}
