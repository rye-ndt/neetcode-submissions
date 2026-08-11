func searchMatrix(m [][]int, target int) bool {
	startR, endR := 0, len(m)-1	

	for startR <= endR {
		midR := (startR + endR) / 2
		row := m[midR]
		l, r := 0, len(row) - 1

		switch {
			case target > row[r]: startR = midR + 1
			case target < row[l]: endR = midR - 1
			default: {
				for l <= r {
					m := (l + r) / 2

					switch {
						case row[m] > target: r = m - 1
						case row[m] < target: l = m + 1
						default: return true 
					}
				}

				return false
			}
		}
	}

	return false
}