type Q struct {
	R int
	C int
}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
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

	time := 0
	directions := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	for len(q) > 0 && fresh > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			rot := q[0]
			q = q[1:]

			for _, dir := range directions {
				nextR := rot.R + dir[0]
				nextC := rot.C + dir[1]
				nextRValid := nextR >= 0 && nextR < rows 
				nextCValid := nextC >= 0 && nextC < cols 

				if nextRValid && nextCValid && grid[nextR][nextC] == 1 {
					fresh--
					grid[nextR][nextC] = 2
					q = append(q, Q{R: nextR, C: nextC})
				}
			}
		}

		time++
	}

	if fresh > 0 {
		return -1
	}

	return time
}
