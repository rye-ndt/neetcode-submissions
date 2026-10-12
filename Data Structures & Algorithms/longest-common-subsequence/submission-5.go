func longestCommonSubsequence(str1 string, str2 string) int {
	memo := map[[2]int]int{}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == len(str1) || j == len(str2) {
			return 0
		}

		key := [2]int{i, j}

		if _, found := memo[key]; found { return memo[key ]}

		if str1[i] == str2[j] {
			memo[key] = 1 + dfs(i+1, j+1)
		} else {
			memo[key] = max(dfs(i, j+1), dfs(i+1, j))
		}

		return memo[key]
	}

	return dfs(0, 0)
}