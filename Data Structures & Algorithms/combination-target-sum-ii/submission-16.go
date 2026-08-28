func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	result := [][]int{}

	var validate func(path []int, index int) bool 
	validate = func(path []int, index int) bool {
		if reduce(path) > target { return false }

		if reduce(path) == target {
			result = append(result, append([]int{}, path...))
			return true 
		}

		lastAccept := -1

		for i := index; i < len(candidates); i++ {
			if candidates[i] == lastAccept { continue }

			if validate(append(path, candidates[i]), i+1) {
				lastAccept = candidates[i]
			}
		}

		return true
	}

	validate([]int{}, 0)

	return result
}

func reduce(v []int) int {
	result := 0
	for _, a := range v {
		result += a
	}
	return result 
}