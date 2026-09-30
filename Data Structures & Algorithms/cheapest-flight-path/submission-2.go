import "slices"

// bellman-ford algorithm 
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	costs := slices.Repeat([]int{math.MaxInt}, n)
	costs[src] = 0 

	for i := 0; i <= k; i++ {
		clone := slices.Clone(costs)

		for _, f := range flights {
			from, to, cost := f[0], f[1], f[2]

			if costs[from] == math.MaxInt { continue }
			clone[to] = min(clone[to], costs[from] + cost)
		}

		costs = clone
	}

	if costs[dst] == math.MaxInt { return -1 }
	return costs[dst]
}