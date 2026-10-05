import "slices"

func coinChange(coins []int, amount int) int {
	dp := slices.Repeat([]int{math.MaxInt}, amount + 1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if c <= i && dp[i - c] < dp[i] {
				dp[i] = dp[i - c] + 1
			}
		}
	}

	if dp[amount] > amount { return -1 }
	
	return dp[amount]
}