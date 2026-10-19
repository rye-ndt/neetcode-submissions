func isNStraightHand(hand []int, groupSize int) bool {
	sort.Ints(hand)
	freq := map[int]int{}

	for _, h := range hand {
		freq[h] += 1
	}

	for _, h := range hand {
		if freq[h] == 0 { continue }

		for j := 0; j < groupSize; j++ {
			if freq[h+j] == 0 { 
				return false 
			}

			freq[h+j]--
		}
	}

	return true 
}
