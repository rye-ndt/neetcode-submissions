func mergeTriplets(triplets [][]int, target []int) bool {
	valids := [][]int{}
	for _, t := range triplets {
		if t[0] <= target[0] && t[1] <= target[1] && t[2] <= target[2] {
			valids = append(valids, t)
		}
	}

	found := [3]int{0, 0, 0}
	for _, v := range valids {
		for i := range v { found[i] = max(found[i], v[i]) }
	}

	for i := range found { 
		if found[i] != target[i] { return false }
	}

	return true
}