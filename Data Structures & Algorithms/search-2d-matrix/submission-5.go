func searchMatrix(m [][]int, target int) bool {
startRow, endRow := 0, len(m)-1	

	for startRow <= endRow {
		midRow := (startRow + endRow) / 2

		cur := m[midRow]

		l, r := 0, len(cur) - 1

		switch {
			case target > cur[r]: startRow = midRow + 1
			case target < cur[l]: endRow = midRow - 1
			default: {
				for l <= r {
					m := (l + r) / 2

					switch {
						case cur[m] > target: r = m - 1
						case cur[m] < target: l = m + 1
						default: return true 
					}
				}

				return false
			}
		}
	}

	return false
}