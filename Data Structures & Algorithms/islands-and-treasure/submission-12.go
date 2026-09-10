const INF = 2147483647

func islandsAndTreasure(grid [][]int) {
	q := [][2]int{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 0 { q = append(q, [2]int{x, y})}
		}	
	}

	dirs := [][2]int{ {1, 0}, {0, 1}, {-1, 0}, {0, -1} }

	for len(q) > 0 {
		lastX, lastY := q[len(q)-1][0], q[len(q)-1][1]

		for _, d := range dirs {
			x, y := lastX + d[0], lastY + d[1]
			validX := x >= 0 && x < len(grid[0])
			validY := y >= 0 && y < len(grid)

			if !validX || !validY || grid[y][x] != INF { continue }
			
			grid[y][x] = grid[lastY][lastX] + 1
			q = append(append([][2]int{}, [2]int{x, y}), q...)
		}
		
		q = q[:len(q)-1]
	}
}