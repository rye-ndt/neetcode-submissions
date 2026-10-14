func findTargetSumWays(nums []int, target int) int {
	cache := map[[2]int]int{}

	var dfs func(i, total int) int 
	dfs = func(i, total int) int {
		k := [2]int{i, total}
		if v, ok := cache[k]; ok { return v }

		if i == len(nums) {
			if total == target { return 1 }
			return 0
		}

		cache[k] = dfs(i+1, total + nums[i]) + dfs(i+1, total - nums[i])
		return cache[k]
	}

	return dfs(0, 0)
}