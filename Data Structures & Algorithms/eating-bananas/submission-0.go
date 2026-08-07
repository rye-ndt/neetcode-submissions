func minEatingSpeed(piles []int, h int) int {
	// core insight: the limit. the max speed must be less than or equal to the max pile 
	minSpeed := 1
	maxSpeed := 0 

	for _, p := range piles {
		if p > maxSpeed {
			maxSpeed = p
		}
	}

	result := maxSpeed

	for minSpeed <= maxSpeed {
		m := (minSpeed + maxSpeed) / 2 // this is different from normal binary search

		finishTime := finishIn(piles, m)

		if finishTime > h {
			// increase the speed 
			minSpeed = m+1
			continue
		}

		if finishTime <= h {
			result = m
			maxSpeed = m-1
			continue
		}
	}

	return result
}

func finishIn(piles []int, speed int) int {
	hours := 0

	for _, p := range piles {
		hours = hours + p / speed

		if p % speed > 0 {
			hours++
		}
	}

	return hours
}
