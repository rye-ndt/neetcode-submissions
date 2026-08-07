func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	sort.Ints(coins)

	best := -1

	for i := len(coins) - 1; i >= 0; i-- {
		c := coins[i]
		if c > amount {
			continue
		}

		// try every count of c, not just the max
		for k := amount / c; k >= 1; k-- {
			sub := coinChange(coins[:i], amount-k*c)
			if sub == -1 {
				continue
			}
			if best == -1 || best > sub+k {
				best = sub + k
			}
		}
	}

	return best
}