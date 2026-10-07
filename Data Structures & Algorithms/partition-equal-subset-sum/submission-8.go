func canPartition(nums []int) bool {
	total := 0
	for _, n := range nums {
		total += n
	}
	if total % 2 != 0 { return false }

	var dfs func(usedIdx []bool, sum int) bool
	dfs = func(usedIdx []bool, sum int) bool {
		if sum == total / 2 { 
			return true
		}

		for i, n := range nums {
			if sum + n > total / 2 || usedIdx[i] { continue }

			usedIdx[i] = true
			if dfs(usedIdx, sum + n) { 
				return true 
			}
			usedIdx[i] = false
		}

		return false
	}	

	return dfs(make([]bool, len(nums)), 0)
}