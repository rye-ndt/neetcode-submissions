func numIslands(grid [][]byte) int {
	result := 0
	directions := [][]int{ {0, 1}, {1, 0}, {0, -1}, {-1, 0} }

	var backtrack func(x, y int) 
	backtrack = func(x, y int) {
		xValid := x >= 0 && x < len(grid[0])
		yValid := y >= 0 && y < len(grid)

		if !xValid || !yValid || grid[y][x] == '0' { return }

		grid[y][x] = '0'

		for _, dir := range directions {
			backtrack(x + dir[0], y + dir[1])
		}
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '0' { continue }
			backtrack(x, y)
			result++
		}
	}

	return result 
}