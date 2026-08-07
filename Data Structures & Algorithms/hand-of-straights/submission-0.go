func isNStraightHand(hand []int, groupSize int) bool {
	sort.Ints(hand)

	fmt.Println("hand: ", hand)

	size := len(hand) / groupSize

	if len(hand) % groupSize > 0 {
		size += 1
	}

	slices := make([][]int, size)

	fmt.Println("len: ", len(slices))

	for i := 0; i < len(hand); i++ {
		cur := hand[i]

		for j := 0; j < len(slices); j++ {
			if slices[j] == nil || len(slices[j]) == 0 {
				slices[j] = []int{cur}
				break
			}

			lastItem := slices[j][len(slices[j]) - 1]

			if cur == lastItem+1 && len(slices[j]) < groupSize {
				slices[j] = append(slices[j], cur)
				break
			}
		}
	}	

	fmt.Println("slices: ", slices)

	for _, s := range slices {
		if len(s) != groupSize {
			return false
		}
	}

	return true
}
