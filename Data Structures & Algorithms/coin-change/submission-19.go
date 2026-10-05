import "slices"

func coinChange(coins []int, amount int) int {
	note := slices.Repeat([]int{math.MaxInt - 1}, amount + 1)
	note[0] = 0

	for i := range note {
		if i == 0 { continue }

		for _, c := range coins {
			if c > i { continue }
			note[i] = min(note[i], note[i - c] + 1)
		}
	}

	if note[amount] > amount { return -1 }
	return note[amount]
}