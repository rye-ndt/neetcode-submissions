func longestCommonSubsequence(str1 string, str2 string) int {
    memo := map[[2]int]int{}

	var solve func(i, j int) int
	solve = func(i, j int) int {
		if i == len(str1) || j == len(str2) {
			return 0
		}

		if _, found := memo[[2]int{i, j}]; found {
			return memo[[2]int{i, j}]
		}

		if str1[i] == str2[j] {
			memo[[2]int{i, j}] = 1 + solve(i+1, j+1)
		} else {
			memo[[2]int{i, j}] = max(solve(i+1, j), solve(i, j+1))
		}

		return memo[[2]int{i, j}]
	}

	return solve(0, 0)
}