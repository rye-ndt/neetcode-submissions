func equals(a, b []int) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2]
}

func mergeTriplets(triplets [][]int, target []int) bool {
	filtered := [][]int{}

	for _, t := range triplets {
		if t[0] <= target[0] && t[1] <= target[1] && t[2] <= target[2] {
			filtered = append(filtered, t)
		}
	}

	fmt.Println("filtered: ", filtered)

	if len(filtered) == 1 {
		return equals(filtered[0], target)
	}

	for i := 1; i < len(filtered); i++ {
		prev := filtered[i-1]
		cur := filtered[i]

		filtered[i] = []int{
			max(prev[0], cur[0]), 
			max(prev[1], cur[1]), 
			max(prev[2], cur[2]), 
		}

		if equals(filtered[i], target) || equals(cur, target) {
			return true
		}
	}

	return false
}
