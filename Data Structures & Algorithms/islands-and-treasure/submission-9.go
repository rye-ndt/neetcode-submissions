const INF = 2147483647

type Q struct {
	R int 
	C int
}

func islandsAndTreasure(grid [][]int) {
	rows := len(grid)
	cols := len(grid[0])
	q := []Q{}

	// push all the treasures into the queue
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				q = append(q, Q{R: r, C: c})
			}
		}
	}	

	directions := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, d := range directions {
			nextR := cur.R + d[0]
			nextC := cur.C + d[1]

			validNextR := nextR >= 0 && nextR < rows
			validNextC := nextC >= 0 && nextC < cols 
			
			if validNextR && validNextC && grid[nextR][nextC] == INF {
				grid[nextR][nextC] = grid[cur.R][cur.C] + 1
				q = append(q, Q{R: nextR, C: nextC})
			}
		}
	}
}
