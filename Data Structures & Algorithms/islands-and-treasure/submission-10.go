const INF = 2147483647

func islandsAndTreasure(grid [][]int) {
	queue := [][2]int{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 0 { queue = append(queue, [2]int{x, y})}
		}	
	}

	dirs := [4][2]int{ {1, 0}, {0, 1}, {-1, 0}, {0, -1} }

	fmt.Println("queue", queue)

	for len(queue) > 0 {
		lastX, lastY := queue[len(queue)-1][0], queue[len(queue)-1][1]

		for _, d := range dirs {
			nextX, nextY := lastX + d[0], lastY + d[1]
			validX := nextX >= 0 && nextX < len(grid[0])
			validY := nextY >= 0 && nextY < len(grid)

			if validX && validY && grid[nextY][nextX] == INF { 
				grid[nextY][nextX] = grid[lastY][lastX] + 1
				queue = append(append([][2]int{}, [2]int{nextX, nextY}), queue...)
			}
		}
		
		queue = queue[:len(queue)-1]
	}
}