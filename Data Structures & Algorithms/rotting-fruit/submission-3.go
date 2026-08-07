type Q struct {
	R int
	C int
}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	time := 0
	q := []Q{}
	fresh := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cur := grid[r][c]

			if cur == 1 {
				fresh++
			}

			if cur == 2 {
				q = append(q, Q{R: r, C: c})
			}
		}
	}

	directions := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	for len(q) > 0 && fresh > 0 {
		thisLayerItems := len(q)
		time++ 

		for i := 0; i < thisLayerItems; i++ {
			cur := q[0]
			q = q[1:]

			for _, d := range directions {
				nextR := cur.R + d[0]
				nextC := cur.C + d[1]
				nextRValid := nextR >= 0 && nextR < rows 
				nextCValid := nextC >= 0 && nextC < cols 

				if nextRValid && nextCValid && grid[nextR][nextC] == 1 {
					grid[nextR][nextC] = 2
					fresh--
					q = append(q, Q{R: nextR, C: nextC})
				}
			}
		}
	}

	if fresh > 0 {
		return -1
	}

	return time
}
