// mental model: 
// which cell can the ocean reach if water flows uphill?

type Q struct {
	R int
	C int
}

func pacificAtlantic(h [][]int) [][]int {
	rows := len(h)
	cols := len(h[0])
	pacific := []Q{}
	atlantic := []Q{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if r == 0 || c == 0 {
				pacific = append(pacific, Q{R: r, C: c})
			}

			if r == rows-1 || c == cols-1 {
				atlantic = append(atlantic, Q{R: r, C: c})
			}
		}
	}

	var directions = [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	var bfs func(sea []Q) map[Q]bool 

	bfs = func (sea []Q) map[Q]bool {
		visited := map[Q]bool{}

		for len(sea) > 0 {
			cur := sea[0]
			visited[cur] = true
			sea = sea[1:]

			for _, dir := range directions {
				nextR := cur.R + dir[0]
				nextC := cur.C + dir[1]
				rValid := nextR >= 0 && nextR < rows
				cValid := nextC >= 0 && nextC < cols 

				if !rValid || !cValid {
					continue
				}

				nextQ := Q{R: nextR, C: nextC}

				if _, found := visited[nextQ]; found {
					continue
				}

				if h[cur.R][cur.C] <= h[nextR][nextC] {
					sea = append(sea, nextQ) 
					visited[nextQ] = true
				}
			}
		}

		return visited
	}

	pacificReached := bfs(pacific)
	atlanticReached := bfs(atlantic)
	result := [][]int{}

	for k, v := range pacificReached {
		if v == false {
			continue
		}
		
		if atlanticReached[k] {
			result = append(result, []int{k.R, k.C})
		}
	}

	return result
}
