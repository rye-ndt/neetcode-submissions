func uniquePaths(m int, n int) int {
	ways := map[[2]int]int{}

	for y := m-1; y >= 0; y-- {
		for x := 0; x < n; x++ {
			if y == m-1 && x == 0 { 
				ways[[2]int{x, y}] = 1 
				continue 
			}

			ways[[2]int{x, y}] = ways[[2]int{x-1, y}] + ways[[2]int{x, y+1}]
		}
	}

	return ways[[2]int{n-1, 0}]
}
