func orangesRotting(grid [][]int) int {
	time := 0

	rotQ := [][2]int{}
	fruits := 0
	dirs := [][2]int{ {1, 0}, {0, 1}, {-1, 0}, {0, -1} }

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 2 { rotQ = append(rotQ, [2]int{x, y}) } 
			if grid[y][x] == 1 { fruits++ }
		}
	}

	for len(rotQ) > 0 {
		rotSomething := false

		for _, rot := range rotQ {
			lastX, lastY := rot[0], rot[1]

			for _, d := range dirs {
				x, y := lastX + d[0], lastY + d[1]
				xValid := x >= 0 && x < len(grid[0])
				yValid := y >= 0 && y < len(grid)

				if !xValid || !yValid || grid[y][x] != 1 { continue }

				grid[y][x] = 2
				fruits--
				rotSomething = true
				rotQ = append(append([][2]int{}, [2]int{x, y}), rotQ...)
			}

			rotQ = rotQ[:len(rotQ)-1]
		}
		
		if rotSomething { time++ }
	}

	if fruits > 0 { return -1 }

	return time
}


/*
120
022
022
*/