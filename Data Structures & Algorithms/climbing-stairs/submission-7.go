func climbStairs(n int) int {
    // algo: ways(n) = ways(n-1) + ways(n-2)
	// so we only have to store 3 vars at once 

	if n <= 2 {
		return n
	}

	// base cases 
	prev1 := 2 // n = 2, 2 ways to climb there 
	prev2 := 1 // n = 1, 1 way to climb there 

	// sliding windows to save space: O(3)
	for i := 3; i <= n; i++ {
		cur := prev1 + prev2
		prev2 = prev1
		prev1 = cur
	}

	return prev1
}
