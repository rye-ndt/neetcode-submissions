// at coin i, you either:
// use it 
// skip it, and never come back 
// -> recursion 
// then you add memoization to speed up 

func change(amount int, coins []int) int {
	memo := map[[2]int]int{}
	
	var count func(i, remain int) int 
	count = func(i, remain int) int {
		if remain == 0 { return 1 }
		if remain < 0 || i == len(coins) { return 0 }

		k := [2]int{i, remain}

		if v, ok := memo[k]; ok { return v }

		use := count(i, remain - coins[i])
		skip := count(i+1, remain)

		memo[k] = use + skip

		return memo[k]
	}

	return count(0, amount)
}
