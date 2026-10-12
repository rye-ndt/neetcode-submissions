func longestCommonSubsequence(a string, b string) int {
	memo := map[[2]int]int{}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == len(a) || j == len(b) {
			return 0
		}

		k := [2]int{i, j}

		if v, ok := memo[k]; ok { return v }

		if a[i] == b[j] {
			memo[k] = 1 + dfs(i+1, j+1)
		} else {
			memo[k] = max(dfs(i, j+1), dfs(i+1, j))
		}

		return memo[k]
	}

	return dfs(0, 0)
}