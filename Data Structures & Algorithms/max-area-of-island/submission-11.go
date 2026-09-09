func maxAreaOfIsland(grid [][]int) int {
	result := 0
	dirs := [][]int{ {1, 0}, {0, 1}, {-1, 0}, {0, -1} }

	var trace func(x, y int) int 
	trace = func(x, y int) int {
		if grid[y][x] == 0 { return 0 }

		grid[y][x] = 0 

		size := 1

		for _, d := range dirs {
			newX := x + d[0]
			newY := y + d[1]

			if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
				size += trace(newX, newY)
			}
		}


		return size
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 { continue }
			result = max(trace(x, y), result)
		}
	}

	return result 
}
