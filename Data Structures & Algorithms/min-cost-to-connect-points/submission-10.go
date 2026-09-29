func minCostConnectPoints(points [][]int) int {
	result := 0
	costs := make([]int, len(points))
	walked := map[int]bool{}

	for i := range costs {
		costs[i] = math.MaxInt
	}

	costs[0] = 0

	for i := 0; i < len(points); i++ {
		nearest := -1 

		for candidate := 0; candidate < len(points); candidate++ {
			if walked[candidate] { continue }

			if nearest == -1 || costs[candidate] < costs[nearest] {
				nearest = candidate
			}
		}

		walked[nearest] = true 
		result += costs[nearest]

		for outside := 0; outside < len(points); outside++ {
			if walked[outside] { continue }
			costs[outside] = min(costs[outside], dist(points[outside], points[nearest]))
		}
	}

	return result
}

func dist(a, b []int) int {
	return max(a[0] - b[0], b[0] - a[0]) + max(a[1] - b[1], b[1] - a[1])
}