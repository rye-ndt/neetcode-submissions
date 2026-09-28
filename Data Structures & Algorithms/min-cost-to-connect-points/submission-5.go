func minCostConnectPoints(points [][]int) int {
	n := len(points)
	total := 0

	onIsland := make([]bool, n)
	bridgeCost := make([]int, n)

	for i := range bridgeCost {
		bridgeCost[i] = math.MaxInt
	}

	bridgeCost[0] = 0 // init 

	for round := 0; round < n; round++ {
		// find the next cheapest point 
		cheapest := -1 

		for cand := 0; cand < n; cand++ {
			if onIsland[cand] { continue }

			if cheapest == -1 ||  bridgeCost[cand] < bridgeCost[cheapest] {
				cheapest = cand
			}
		}
		
		// add it
		onIsland[cheapest] = true 
		total += bridgeCost[cheapest]
		
		// find the new cheapest to all outsiders
		for outside := 0; outside < n; outside++ {
			if onIsland[outside] { continue }
			bridgeCost[outside] = min(bridgeCost[outside], dist(points[cheapest], points[outside]))
		}
	}


	return total
}

func dist(a, b []int) int {
	abs := func(x int) int {
		if x < 0 { return -x }
		return x
	}

	return abs(a[0] - b[0]) + abs(a[1] - b[1])
}