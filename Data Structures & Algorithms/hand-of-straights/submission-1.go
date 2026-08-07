func isNStraightHand(hand []int, groupSize int) bool {
	sort.Ints(hand)

	size := len(hand) / groupSize

	if len(hand) % groupSize > 0 {
		size += 1
	}

	slices := make([][]int, size)

	for i := 0; i < len(hand); i++ {
		for j := 0; j < len(slices); j++ {
			if slices[j] == nil || len(slices[j]) == 0 {
				slices[j] = []int{hand[i]}
				break
			} else if hand[i] == slices[j][len(slices[j]) - 1]+1 && len(slices[j]) < groupSize {
				slices[j] = append(slices[j], hand[i])
				break
			}
		}
	}	

	for _, s := range slices {
		if len(s) != groupSize {
			return false
		}
	}

	return true
}
