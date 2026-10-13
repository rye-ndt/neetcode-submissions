// at coin i, you either:
// use it 
// skip it, and never come back 
// -> recursion 
// then you add memoization to speed up 

func change(amount int, coins []int) int {
	memo := map[[2]int]int{}

	var dfs func(i, remain int) int 
	dfs = func(i, remain int) int {
		k := [2]int{i, remain}
		if v, ok := memo[k]; ok { return v }

		switch {
			case i == len(coins): return 0
			case remain < 0: return 0
			case remain == 0: return 1
			default:
				use := dfs(i, remain - coins[i])
				skip := dfs(i+1, remain)

				memo[k] = use + skip

				return memo[k]
		}
	}

	return dfs(0, amount)
}