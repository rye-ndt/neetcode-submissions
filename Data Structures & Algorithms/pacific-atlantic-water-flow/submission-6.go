func pacificAtlantic(h [][]int) [][]int {
	result := [][]int{}
	aReach, pReach := map[[2]int]bool{}, map[[2]int]bool{}
	dirs := [][2]int{ {0, 1}, {1, 0}, {0, -1}, {-1, 0} }

	var dfs func(x, y int, seen map[[2]int]bool)
	dfs = func(x, y int, seen map[[2]int]bool) {
		if seen[[2]int{y, x}] { return }
		seen[[2]int{y, x}] = true

		for _, d := range dirs {
			nextX, nextY := x + d[0], y + d[1]
			xValid := nextX >= 0 && nextX < len(h[0])
			yValid := nextY >= 0 && nextY < len(h)

			if xValid && yValid && h[nextY][nextX] >= h[y][x] {
				dfs(nextX, nextY, seen)
			}
		}
	}

	for x := 0; x < len(h[0]); x++ {
		for y := 0; y < len(h); y++ {
			if x == 0 || y == 0 { dfs(x, y, pReach) }
			if x == len(h[0])-1 || y == len(h)-1 { dfs(x, y, aReach) }
		}
	}

	for k, _ := range aReach {
		if pReach[k] { result = append(result, []int{k[0], k[1]}) }
	}

	return result
}