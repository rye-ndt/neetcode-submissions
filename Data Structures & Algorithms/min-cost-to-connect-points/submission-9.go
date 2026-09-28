func minCostConnectPoints(points [][]int) int {
	bridgeCost := make([]int, len(points))
	walked := make([]bool, len(points))
	total := 0

	for i := range bridgeCost { bridgeCost[i] = 1_000_000_000 }
	bridgeCost[0] = 0 // to init the *, or will stuck 

	for i := 0; i < len(points); i++ {
		cheapest := -1 

		for can := 0; can < len(points); can++ { // *
			if walked[can] { continue }
			if cheapest == -1 || bridgeCost[can] < bridgeCost[cheapest] {
				cheapest = can
			}
		}

		walked[cheapest] = true
		total += bridgeCost[cheapest]

		// update other points
		for outside := 0; outside < len(points); outside++ {
			if walked[outside] { continue }
			bridgeCost[outside] = min(bridgeCost[outside], dist(points[cheapest], points[outside]))
		}
	}

	return total
}

func dist(a, b []int) int {
	return max(a[0] - b[0], b[0] - a[0]) + max(a[1] - b[1], b[1] - a[1])
}