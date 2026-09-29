import "slices"

// bellman-ford algorithm 
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cost := slices.Repeat([]int{math.MaxInt}, n)
	cost[src] = 0

	for i := 0; i <= k; i++ {
		today := slices.Clone(cost)

		// try all flights 
		for _, f := range flights {
			from, to, price := f[0], f[1], f[2]

			if cost[from] == math.MaxInt { continue }
			today[to] = min(today[to], cost[from] + price)
		}

		cost = today
	}

	if cost[dst] == math.MaxInt { return -1 }
	return cost[dst]
}