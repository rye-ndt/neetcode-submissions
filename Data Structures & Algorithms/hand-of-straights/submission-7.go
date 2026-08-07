// brute force O(n^2) done 

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand) % groupSize != 0 {
		return false
	}

	store := map[int]int{}
	for _, n := range hand {
		store[n]++
	}

	sort.Ints(hand)

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
