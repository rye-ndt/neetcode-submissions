func minCostConnectPoints(points [][]int) int {
	bridgeCost := make([]int, len(points))
	walked := make([]bool, len(points))
	total := 0

	for i := range bridgeCost { bridgeCost[i] = 1_000_000_000 }
	bridgeCost[0] = 0 // to init the *

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

		for outside := 0; outside < len(points); outside++ {
			if walked[outside] { continue }
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