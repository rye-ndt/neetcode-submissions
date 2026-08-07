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

	fmt.Println("start: ", nums)

	i := 0

	for i < len(nums) {
		fmt.Println("considering: ", nums[i])
		j := i+1

		for j < len(nums) {
			fmt.Println("merging with: ", nums[j])

			if !overlap(nums[i], nums[j]) {
				break
			}

			nums[i] = join(nums[i], nums[j])
			nums[j] = []int{-1, -1}
			fmt.Println("after merged: ", nums[i])

			j++
		}

		i = j
		fmt.Println("now i is: ", i)
	}

	result := [][]int{}

	for _, n := range nums {
		if n[0] != -1 || n[1] != -1 {
			result = append(result, n)
		}
	}

	return result
}
