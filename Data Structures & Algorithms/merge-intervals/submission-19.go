func merge(nums [][]int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	clone := [][]int{nums[0]}

	for i := 1; i < len(nums); i++ {
		cur, last := nums[i], clone[len(clone)-1]

		if cur[0] > last[1] || cur[1] < last[0] {
			clone = append(clone, cur)
			continue 
		}

		clone = append(clone[:len(clone)-1], []int{min(cur[0], last[0]), max(cur[1], last[1])})
	}

	return clone
}