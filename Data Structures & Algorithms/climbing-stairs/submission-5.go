func climbStairs(n int) int {
    if n <= 2 {
		return n
	}

	prev2 := 1
	prev1 := 2

	for i := 3; i <=n; i++ {
		cur := prev1 + prev2
		prev2 = prev1
		prev1 = cur
	}

	return prev1
}
