func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand) % groupSize > 0 {
		return false 
	}
	
	count := map[int]int{}

	for _, c := range hand {
		count[c]++
	}

	sort.Ints(hand)

	for _, c := range hand {
		if count[c] == 0 {
			continue
		}

		for i := c; i < c + groupSize; i++ {
			if count[i] == 0 {
				return false
			}

			count[i]--
		}
	}

	return true
}
