func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	result := [][]int{}

	var validate func(path []int, index int) bool 
	validate = func(path []int, index int) bool {
		if reduce(path) == target {
			result = append(result, append([]int{}, path...))
			return true 
		}

		if reduce(path) > target { return false }

		lastAccept := -1

		for i := index; i < len(candidates); i++ {
			if candidates[i] == lastAccept { continue }

			stored := validate(append(path, candidates[i]), i+1) 

			if stored {
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