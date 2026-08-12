func minEatingSpeed(piles []int, time int) int {
	l, r := 1, piles[0]

	for _, p := range piles {
		r = max(r, p)
	}	

	lowest := r

	for l <= r {
		m := (r + l) / 2

		ateIn := 0

		for _, p := range piles {
			ateIn += (p + m - 1) / m
		}

		fmt.Println("r, l", r, l, "with m", m, "ate in", ateIn, "current lowest", lowest)

		if ateIn <= time {
			lowest = min(lowest, m)
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return lowest
}