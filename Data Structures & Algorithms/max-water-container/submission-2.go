func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func maxArea(h []int) int {	
	result := 0

	l, r := 0, len(h) - 1

	for l < r {
		w := r - l 

		size := w * min(h[l], h[r])

		if size > result {
			result = size
		}

		if h[l] > h[r] {
			r--
			continue
		}

		l++
	}

	return result
}
