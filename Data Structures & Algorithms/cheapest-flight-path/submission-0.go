import "slices"

// bellman-ford algorithm 
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	best := make([]int, n)
	for i := range best {
		best[i] = math.MaxInt
	}

	best[src] = 0

	for i := 0; i <= k; i++ {
		today := slices.Clone(best)

		for _, f := range flights {
			from, to, price := f[0], f[1], f[2]

			if best[from] == math.MaxInt { continue }	

			today[to] = min(today[to], best[from] + price)
		}

		best = today
	}

	if best[dst] == math.MaxInt { return -1 }

	return best[dst]	
}