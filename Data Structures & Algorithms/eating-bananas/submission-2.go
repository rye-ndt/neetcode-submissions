// core insight: the limit. the max speed must be less than or equal to the max pile 
// so you will find the minimum value satisfies the question, in the range 1 < x < max pile
// so now you binary search the result

func minEatingSpeed(piles []int, h int) int {
	min := 1
	max := 0

	for _, p := range piles {
		if p > max {
			max = p
		}
	}

	minimumResult := max

	for min <= max {
		mid := min + (max - min) / 2 

		finishTime := finishIn(piles, mid)

		if finishTime > h {
			// increase the speed
			min = mid + 1
			continue
		}

		minimumResult = mid

		// continue, since we might find something smaller ahead 
		max = mid-1
	}

	return minimumResult
}

func finishIn(piles []int, speed int) int {
	h := 0

	for _, p := range piles {
		h += (p / speed)

		if p % speed > 0 {
			h++
		}
	}

	return h
}
