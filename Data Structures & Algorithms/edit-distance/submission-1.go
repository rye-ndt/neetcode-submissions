func minDistance(w1 string, w2 string) int {
	memo := map[[2]int]int{}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == len(w1) { return len(w2) - j }
		if j == len(w2) { return len(w1) - i }

		k := [2]int{i, j}
		if v, ok := memo[k]; ok { return v }

		memo[k] = math.MaxInt 

		if w1[i] == w2[j] {
			memo[k] = min(memo[k], dfs(i+1, j+1))
		}

		if w1[i] != w2[j] {
			memo[k] = min(memo[k], 1 + dfs(i, j+1))
			memo[k] = min(memo[k], 1 + dfs(i+1, j))
			memo[k] = min(memo[k], 1 + dfs(i+1, j+1))
		}


		return memo[k]
	}

	return dfs(0, 0)
}
