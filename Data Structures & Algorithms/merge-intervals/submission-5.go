func join(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func overlap(a, b []int) bool {
	return b[0] <= a[1]
}

func merge(nums [][]int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	result := [][]int{}

	i := 0

	for i < len(nums) {
		j := i+1

		for j < len(nums) {
			if !overlap(nums[i], nums[j]) {
				break
			}

			nums[i] = join(nums[i], nums[j])

			j++
		}

		result = append(result, nums[i])

		i = j
	}

	return result
}
