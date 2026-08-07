func min(a, b int) int {
	if a > b {
		return b
	}	

	return a
}	

func maxArea(heights []int) int {
	max := 0

	l := 0
	r := len(heights) - 1

	for l < r {
		smaller := min(heights[l], heights[r])

		square := (r - l) * smaller

		if square > max {
			max = square 
		}

		if smaller == heights[l] {
			l++
			continue
		}

		r--
	}

	return max
}
