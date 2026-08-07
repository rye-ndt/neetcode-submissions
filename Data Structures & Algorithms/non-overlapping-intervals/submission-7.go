func overlap(a, b []int) bool {
	return b[0] < a[1]
}

func merge(a, b []int) []int {
	return []int{min(a[0], b[0]), min(a[1], b[1])}
}

func eraseOverlapIntervals(nums [][]int) int {
    sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	fmt.Println("start: ", nums)

	i := 0
	counter := 1

	for i < len(nums) {
		j := i+1

		for j < len(nums) {
			if !overlap(nums[i], nums[j]) {
				counter++
				break
			}

			nums[i] = merge(nums[i], nums[j])			
			j++
		}

		i = j
	}

	return len(nums) - counter
}
