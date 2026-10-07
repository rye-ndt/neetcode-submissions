func canPartition(nums []int) bool {
	total := 0
	for _, n := range nums {
		total += n
	}
	if total % 2 != 0 { return false }

	target := total / 2
	can := false 
	sort.Ints(nums)

	var dfs func(usedIndices []bool, sum int) 
	dfs = func(usedIndices []bool, sum int) {
		if sum == target || can { 
			can = true 
			return 
		}

		for i, n := range nums {
			if sum + n > target || usedIndices[i] { continue }

			usedIndices[i] = true
			dfs(usedIndices, sum + n)
			usedIndices[i] = false
		}
	}	

	dfs(make([]bool, len(nums)), 0)

	return can
}