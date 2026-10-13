func findTargetSumWays(nums []int, target int) int {
	memo := map[[2]int]int{}

	var dfs func(i, sum int) int 
	dfs = func(i, sum int) int {
		k := [2]int{i, sum}
		if v, ok := memo[k]; ok { return v }

		if i == len(nums) { 
			if sum == target { return 1 }
			return 0
		}

		memo[k] = dfs(i+1, sum + nums[i]) + dfs(i+1, sum - nums[i])

		return memo[k]
	}	

	return dfs(0, 0)
}