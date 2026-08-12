func minEatingSpeed(piles []int, time int) int {
	l, r := 1, piles[0] 

	for _, p := range piles {
		r = max(r, p)
	}

	result := r

	for l <= r {
		m := (r + l) / 2

		ateIn := 0
		for _, p := range piles {
			ateIn += (p + m - 1) / m
		}

		switch {
			case ateIn <= time:
				result = min(result, m)
				r = m - 1 
			default: l = m + 1
		}
	}

	return result
}