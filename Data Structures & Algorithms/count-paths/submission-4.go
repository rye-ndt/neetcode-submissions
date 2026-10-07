import "slices"

func uniquePaths(m int, n int) int {
	dp := slices.Repeat([]int{1}, n)

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}

	return dp[n-1]
}
