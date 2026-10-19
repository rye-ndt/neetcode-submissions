func isNStraightHand(hand []int, groupSize int) bool {
	sort.Ints(hand)

	f := map[int]int{}
	for _, num := range hand { f[num]++ }

	for _, num := range hand {
		if f[num] == 0 { continue }

		for i := 0; i < groupSize; i++ {
			if f[num + i] == 0 { return false }
			f[num + i]--
		}
	}

	return true 
}