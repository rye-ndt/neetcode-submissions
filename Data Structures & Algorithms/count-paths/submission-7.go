func uniquePaths(m int, n int) int {
	ways := map[[2]int]int{}
	ways[[2]int{0, 0}] = 1

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			k := [2]int{x, y}
			if _, ok := ways[k]; ok { continue }
			
			ways[k] = ways[[2]int{x, y-1}] + ways[[2]int{x-1, y}]
		}
	}

	return ways[[2]int{n-1, m-1}]
}
