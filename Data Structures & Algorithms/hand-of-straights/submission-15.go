func isNStraightHand(cards []int, size int) bool {
	sort.Ints(cards)

	f := map[int]int{}
	for _, c := range cards { f[c]++ }

	for _, c := range cards {
		if f[c] == 0 { continue }

		for i := 0; i < size; i++ {
			if f[c + i] == 0 { return false }
			f[c + i]--
		}
	}

	return true 
}