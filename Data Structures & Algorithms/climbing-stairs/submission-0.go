func climbStairs(n int) int {
    if n <= 2 {
		return n
	}

	notebook := make([]int, n+1)

	// base cases
	notebook[1] = 1
	notebook[2] = 2 

	for i := 3; i <= n; i++ {
		notebook[i] = notebook[i-1] + notebook[i-2]
	}

	return notebook[n]
}
