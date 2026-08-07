func isNStraightHand(hand []int, groupSize int) bool {
	sort.Ints(hand)

	store := map[int]int{}

	for _, n := range hand {
		store[n]++
	}

	for _, n := range hand {
		if store[n] == 0 {
			continue
		}

		for i := n; i < n + groupSize; i++ {
			if store[i] == 0 {
				return false
			}

			store[i]--
		}
	}

	return true
}
