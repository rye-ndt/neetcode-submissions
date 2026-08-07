const INF = 2147483647

func islandsAndTreasure(grid [][]int) {
	rows := len(grid)
	cols := len(grid[0])

	type cell struct {
		r int
		c int
	}

	q := []cell{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				q = append(q, cell{r, c})
			}
		}
	}

	dirs := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, d := range dirs {
			nr := cur.r + d[0]
			nc := cur.c + d[1]

			nrValid := nr >= 0 && nr < rows
			ncValid := nc >= 0 && nc < cols 

			if nrValid && ncValid && grid[nr][nc] == INF {
				grid[nr][nc] = grid[cur.r][cur.c] + 1
				q = append(q, cell{nr, nc})
			}
		}
	}
}
