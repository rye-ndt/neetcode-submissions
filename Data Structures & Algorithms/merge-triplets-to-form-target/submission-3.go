func mergeTriplets(triplets [][]int, target []int) bool {
	// filter 
	valids := [][]int{}

	for _, t := range triplets {
		if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] {
			continue 
		}

		valids = append(valids, t)
	}

	found := [3]int{0, 0, 0}

	for _, v := range valids {
		for i, item := range v { found[i] = max(found[i], item) }
	}

	for i := range found { 
		if found[i] != target[i] { return false }
	}

	return true
}