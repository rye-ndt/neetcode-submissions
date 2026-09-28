func minCostConnectPoints(points [][]int) int {
	n := len(points)
	dist := make([]int, n)
	inTree := make([]bool, n)

	for i := range dist {
		if i > 0 { dist[i] = 1_000_000_000 }
	}

	cost := 0

	for k := 0; k < n; k++ {
		u := -1
		for i := 0; i < n; i++ {
			if !inTree[i] && (u == -1 || dist[i] < dist[u]) {
				u = i
			}
		}
		inTree[u] = true
		cost += dist[u]

		for v := 0; v < n; v++ {
			if !inTree[v] {
				d := abs(points[u][0] - points[v][0]) + abs(points[u][1] - points[v][1]) 
				if d < dist[v] { dist[v] = d }
			}
		}
	}

	return cost
}

func abs(x int) int { 
	if x > 0 { return x }
	return -x
}