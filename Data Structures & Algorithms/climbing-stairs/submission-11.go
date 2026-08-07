func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	
	note := make([]int, n)
	note[0] = 1
	note[1] = 2 

	for i := 2; i < n; i++ {
		note[i] = note[i-1] + note[i-2]
	}

	return note[n-1]
}
