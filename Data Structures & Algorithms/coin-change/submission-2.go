// at coin i 
// try add it up multiple times, until it meet the amount or surpass

func coinChange(coins []int, amount int) int {
	cannot := math.MaxInt32 

	note := make([]int, amount+1)

	for i := range note {
		note[i] = cannot
	}

	note[0] = 0

	for x := 1; x <= amount; x++ {
		for _, c := range coins {
			if c <= x && note[x-c] != cannot {
				note[x] = min(note[x], 1+note[x-c])
			}
		}
	}

	if note[amount] == cannot {
		return -1
	}
	
	return note[amount]
}
