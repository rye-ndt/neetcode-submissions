func longestCommonSubsequence(text1 string, text2 string) int {
    memo := map[[2]int]int{}

	var solve func(i, j int) int
	solve = func(i, j int) int {
		if i == len(text1) || j == len(text2) {
			return 0
		}

		if _, found := memo[[2]int{i, j}]; found {
			return memo[[2]int{i, j}]
		}

		result := 0

		if text1[i] == text2[j] {
			result = 1 + solve(i+1, j+1)
		} else {
			result = max(solve(i+1, j), solve(i, j+1))
		}

		memo[[2]int{i, j}] = result
		return result
	}

	return solve(0, 0)
}
