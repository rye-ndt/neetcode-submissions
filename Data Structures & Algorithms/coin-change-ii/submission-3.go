// at coin i, you either:
// use it 
// skip it, and never come back 
// -> recursion 
// then you add memoization to speed up 

func change(amount int, coins []int) int {
	memo := map[[2]int]int{}

	var dfs func(i, remain int) int 
	dfs = func(i, remain int) int {
		if i >= len(coins) { return 0 }
		if remain < 0 { return 0 }
		if remain == 0 { return 1 }

		k := [2]int{i, remain}

		if v, ok := memo[k]; ok { return v }

		use := dfs(i, remain - coins[i])
		skip := dfs(i+1, remain)

		memo[k] = use + skip

		return use + skip
	}

	return dfs(0, amount)
}
